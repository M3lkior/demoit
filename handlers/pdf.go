/*
Copyright 2018 Google LLC
Copyright 2022 David Gageot

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package handlers

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/dgageot/demoit/files"
	"github.com/dgageot/demoit/flags"
	"github.com/jung-kurt/gofpdf"
)

const (
	width  = 1920.0
	height = 1080.0
	zoom   = 2.0
)

type image struct {
	buf []byte
	err error
}

// ExportToPDF generates a pdf that contains one page per slide.
func ExportToPDF(w http.ResponseWriter, r *http.Request) {
	if err := exportPagesToPdf(r.Context(), w); err != nil {
		http.Error(w, fmt.Sprintf("Unable export to pdf: %v", err), http.StatusInternalServerError)
		return
	}
}

func exportPagesToPdf(ctx context.Context, w http.ResponseWriter) error {
	pageCount, err := readPageCount()
	if err != nil {
		return err
	}

	var images [](chan image)
	for p := 0; p < pageCount; p++ {
		images = append(images, make(chan image))
	}

	readPagesAsPNG(ctx, images)

	pdf, err := writePagesToPDF(images)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/pdf")
	return pdf.Output(w)
}

func readPagesAsPNG(ctx context.Context, images [](chan image)) {
	var actions [4][]chromedp.Action

	group := 0
	for i, result := range images {
		i := i
		result := result

		actions[group] = append(actions[group],
			// The viewport is sized before the slide loads, not after: the
			// stage computes its scale on load, and mermaid lays its diagram
			// out against that scale. Resized afterwards, the capture caught a
			// diagram laid out for another viewport and not yet repainted --
			// an empty pane.
			chromedp.ActionFunc(func(ctx context.Context) error {
				if err := emulation.SetDeviceMetricsOverride(width*int64(zoom), height*int64(zoom), 1, false).Do(ctx); err != nil {
					result <- image{err: err}
					return err
				}
				return nil
			}),
			// theme=light is forced rather than inherited: a PDF of dark slides
			// is a PDF of ink. The query parameter wins over whatever the
			// browser remembered, the same way ?grid=true already steers a
			// render. export=pdf shows every progressive-reveal step, which a
			// single capture has nobody to walk, and hides the GitHub star
			// button, a live count that has no business on paper.
			chromedp.Navigate(fmt.Sprintf("http://%s/%d?theme=light&display=screen&export=pdf", flags.WebServerAddress(), i)),
			waitForMermaid(),
			chromedp.ActionFunc(func(ctx context.Context) error {
				buf, err := page.CaptureScreenshot().WithClip(&page.Viewport{
					Width:  width * zoom,
					Height: height * zoom,
					Scale:  1,
				}).Do(ctx)
				if err != nil {
					result <- image{err: err}
					return err
				}

				fmt.Println("Exported page", i)
				result <- image{buf: buf}
				return nil
			}))

		group = (group + 1) % 4
	}

	for _, tasks := range actions {
		go func(tasks []chromedp.Action) {
			ctx, cancel := chromedp.NewContext(ctx, chromedp.WithErrorf(logBrowserError))
			defer cancel()

			chromedp.Run(ctx, tasks...)
		}(tasks)
	}
}

// mermaidRendered is true once every diagram on the slide is drawn, and
// trivially true on a slide without one. data-processed alone is not enough:
// mermaid sets it before it renders, not after. Nor is the SVG: the talk's
// renderMermaid counter-scales the container while mermaid measures, and only
// removes that inline transform once mermaid.run has settled -- a capture
// taken in between shows the diagram blown up out of its pane.
const mermaidRendered = `Array.from(document.querySelectorAll('.mermaid')).every(n => n.dataset.processed === 'true' && n.querySelector('svg') && !n.style.transform)`

// waitForMermaid holds the capture until the slide's diagrams are drawn.
// Navigate returns on the load event, and mermaid renders later than that: it
// waits for the fonts first, then lays out asynchronously, and the talk keeps
// the element hidden until it is done -- so a capture taken on load prints an
// empty pane. A diagram that never finishes (a syntax error, no network for
// the CDN) must not cost the whole export, so a timeout captures the slide as
// it stands rather than failing.
func waitForMermaid() chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		var done bool
		if err := chromedp.Poll(mermaidRendered, &done, chromedp.WithPollingTimeout(10*time.Second)).Do(ctx); err != nil {
			fmt.Println("Mermaid not rendered before capture:", err)
		}

		// Two frames, so what the condition saw has been painted. Not fatal
		// either: an error returned here would stop the tab before the capture
		// action ever reports to its channel, and hang the whole export.
		if err := chromedp.Evaluate(`new Promise(r => requestAnimationFrame(() => requestAnimationFrame(() => r(true))))`, &done, func(p *runtime.EvaluateParams) *runtime.EvaluateParams {
			return p.WithAwaitPromise(true)
		}).Do(ctx); err != nil {
			fmt.Println("Could not wait for a paint before capture:", err)
		}
		return nil
	})
}

// logBrowserError drops the events the vendored cdproto cannot decode. It is
// older than the Chrome it drives, so every enum value Chrome added since
// (an IPAddressSpace of "Loopback", say) fails to unmarshal and chromedp logs
// it as an error -- on every request, while the export works perfectly well.
// Those events are ones this code never listens to.
func logBrowserError(format string, args ...interface{}) {
	if strings.HasPrefix(format, "could not unmarshal event") {
		return
	}
	log.Printf("ERROR: "+format, args...)
}

func writePagesToPDF(images [](chan image)) (*gofpdf.Fpdf, error) {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr: "cm",
		Size:    gofpdf.SizeType{Wd: 29.7, Ht: 29.7 * height / width},
	})

	for i, result := range images {
		image := <-result
		if image.err != nil {
			return nil, image.err
		}

		imageName := fmt.Sprintf("image%d", i)
		pdf.AddPage()
		pdf.RegisterImageReader(imageName, "png", bytes.NewReader(image.buf))
		pdf.ImageOptions(imageName, 0, 0, 29.7, 0, false, gofpdf.ImageOptions{ImageType: "png", ReadDpi: true}, 0, "")

		if err := pdf.Error(); err != nil {
			return nil, err
		}
	}

	return pdf, nil
}

func readPageCount() (int, error) {
	steps, err := readSteps(files.Root)
	if err != nil {
		return 0, err
	}

	return len(steps), nil
}
