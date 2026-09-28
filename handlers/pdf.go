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
	"context"
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/dgageot/demoit/files"
	"github.com/dgageot/demoit/flags"
)

//go:embed resources/print.tmpl.html
var printHTML string
var printTemplate = template.Must(template.New("print").Funcs(templateFuncs).Parse(printHTML))

// Print renders every step into one printable page, one step per page, for
// ExportToPDF to hand to Chrome's own PDF printer.
func Print(w http.ResponseWriter, r *http.Request) {
	steps, err := readSteps(files.Root)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to read steps: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	if err := printTemplate.Execute(w, steps); err != nil {
		http.Error(w, "Unable to render print view", http.StatusInternalServerError)
		return
	}
}

// ExportToPDF generates a pdf that contains one page per slide.
//
// Chrome prints /print, the whole deck in one document, with its own PDF
// printer rather than having each slide screenshotted: text and shapes stay
// vectors, and fonts are embedded once for the deck. The screenshots it
// replaces made a 42-slide deck weigh 34MB, one 3840x2160 bitmap a page, where
// this one is under 1MB.
func ExportToPDF(w http.ResponseWriter, r *http.Request) {
	buf, err := printDeck(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable export to pdf: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Write(buf)
}

func printDeck(ctx context.Context) ([]byte, error) {
	ctx, cancel := chromedp.NewContext(ctx, chromedp.WithErrorf(logBrowserError))
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(fmt.Sprintf("http://%s/print", flags.WebServerAddress())),
		waitForSlides(),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			// The page size comes from the @page rule of print.tmpl.html,
			// which is the stage's own 1920x1080.
			buf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithPreferCSSPageSize(true).
				WithMarginTop(0).
				WithMarginBottom(0).
				WithMarginLeft(0).
				WithMarginRight(0).
				Do(ctx)
			return err
		}),
	)

	return buf, err
}

// slidesReady is true once every slide's iframe has loaded and every diagram
// in it is drawn. data-processed alone is not enough: mermaid sets it before
// it renders, not after. Nor is the SVG: the talk's renderMermaid
// counter-scales the container while mermaid measures, and only removes that
// inline transform once mermaid.run has settled -- a print taken in between
// shows the diagram blown up out of its pane, or not at all.
const slidesReady = `Array.from(document.querySelectorAll('iframe')).every(f => {
	const d = f.contentDocument;
	return d && d.readyState === 'complete' && d.location.href !== 'about:blank' &&
		Array.from(d.querySelectorAll('.mermaid')).every(n => n.dataset.processed === 'true' && n.querySelector('svg') && !n.style.transform);
})`

// waitForSlides holds the print until every slide is ready. The load event
// of /print comes before mermaid is done: it waits for the fonts first, then
// lays out asynchronously, and the talk keeps the element hidden until then.
// A slide that never finishes (a diagram with a syntax error, no network for
// the CDN) must not cost the whole export, so a timeout prints the deck as it
// stands rather than failing.
func waitForSlides() chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		var done bool
		if err := chromedp.Poll(slidesReady, &done, chromedp.WithPollingTimeout(30*time.Second)).Do(ctx); err != nil {
			fmt.Println("Slides not all rendered before print:", err)
		}

		// Two frames, so what the condition saw has been laid out.
		if err := chromedp.Evaluate(`new Promise(r => requestAnimationFrame(() => requestAnimationFrame(() => r(true))))`, &done, func(p *runtime.EvaluateParams) *runtime.EvaluateParams {
			return p.WithAwaitPromise(true)
		}).Do(ctx); err != nil {
			fmt.Println("Could not wait for a frame before print:", err)
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
