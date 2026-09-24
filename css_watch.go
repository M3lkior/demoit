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

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// cssScript is the stylesheet builder, as it is reachable from the repository
// root. A released binary run from anywhere else will not find it, which is
// the case stylesheetBuilder reports rather than fails on.
const cssScript = "hack/css.sh"

// stylesheetOutput is the file the builder writes, and therefore the one file
// change that must never trigger another build.
const stylesheetOutput = "tailwind.css"

// stylesheetBuilder rebuilds a talk's Tailwind stylesheet when one of its
// files changes, so that a class written in a slide exists by the time the
// browser reloads. Without it the class silently does not exist: nothing in
// the Go build regenerates these stylesheets, and a browser does not report an
// unknown class — the slide just renders without the style that was asked for.
//
// It drives the one-shot build on demoit's own file watcher rather than
// running `hack/css.sh --watch` beside the server. The CLI's own watch mode
// does not see these files: the sources come from an `@source "../"` in
// tailwind.src.css rather than from the entry point, and a change to a slide
// never rebuilt anything. demoit is already watching the presentation folder
// for live reload, so this reuses the signal that does work, and the build is
// a few tens of milliseconds.
type stylesheetBuilder struct {
	folder string
}

// newStylesheetBuilder returns the builder for a talk, or nil and the reason
// there is none.
func newStylesheetBuilder(folder string) (*stylesheetBuilder, string) {
	if _, err := os.Stat(cssScript); err != nil {
		return nil, fmt.Sprintf("%s is not here, so this is not a run from the repository. Rebuild the stylesheet by hand after writing a new Tailwind class.", cssScript)
	}

	// A talk without its own build entry does not use Tailwind at all --
	// sample/ predates the chassis and leans on its style.css.
	source := filepath.Join(folder, ".demoit", "tailwind.src.css")
	if _, err := os.Stat(source); err != nil {
		return nil, fmt.Sprintf("%s has no %s, so it has no Tailwind build of its own.", folder, source)
	}

	return &stylesheetBuilder{folder: folder}, ""
}

// rebuild runs the builder for the change at path and reports whether it ran.
func (b *stylesheetBuilder) rebuild(path string) bool {
	if b == nil || !triggersRebuild(path) {
		return false
	}

	out, err := exec.Command(cssScript, b.folder).CombinedOutput()
	if err != nil {
		// A broken stylesheet is worth saying out loud and worth carrying on
		// from: the previous build is still on disk, so the deck keeps the
		// styling it had rather than losing all of it over one bad class.
		fmt.Printf("Stylesheet build failed: %v\n%s", err, out)
	}

	return true
}

// triggersRebuild reports whether a changed file can change the stylesheet.
//
// The output is excluded first, and that is not an optimisation: the builder
// writes it inside the folder demoit is watching, so treating it as a source
// would have every build trigger the next one.
func triggersRebuild(path string) bool {
	if filepath.Base(path) == stylesheetOutput {
		return false
	}

	switch strings.ToLower(filepath.Ext(path)) {
	case ".md", ".html", ".css", ".yml", ".yaml":
		return true
	default:
		return false
	}
}
