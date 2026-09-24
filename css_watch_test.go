package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// chdir moves into dir for the duration of the test. The decision reads two
// paths relative to the working directory, which is the whole point: the
// script only exists in a checkout.
func chdir(t *testing.T, dir string) {
	t.Helper()

	previous, err := os.Getwd()
	if err != nil {
		t.Fatalf("unable to read the working directory: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("unable to enter %s: %v", dir, err)
	}
	t.Cleanup(func() { _ = os.Chdir(previous) })
}

func TestStylesheetBuilderRunsForATalkThatHasATailwindBuild(t *testing.T) {
	root := t.TempDir()
	write(t, filepath.Join(root, cssScript), "#!/bin/sh\n")
	write(t, filepath.Join(root, "talk", ".demoit", "tailwind.src.css"), "")
	chdir(t, root)

	builder, reason := newStylesheetBuilder("talk")
	if builder == nil {
		t.Fatalf("got no builder (%s), want one", reason)
	}
	if builder.folder != "talk" {
		t.Errorf("got folder %q, want the talk being served", builder.folder)
	}
}

func TestStylesheetBuilderStandsDownOutsideACheckout(t *testing.T) {
	// A released binary run from anywhere: no script to run, and no reason to
	// stop the server over it.
	root := t.TempDir()
	write(t, filepath.Join(root, "talk", ".demoit", "tailwind.src.css"), "")
	chdir(t, root)

	builder, reason := newStylesheetBuilder("talk")
	if builder != nil {
		t.Fatalf("got a builder, want none when %s is absent", cssScript)
	}
	if !strings.Contains(reason, cssScript) {
		t.Errorf("got reason %q, want it to name the missing script", reason)
	}
}

func TestStylesheetBuilderStandsDownForATalkWithoutTailwind(t *testing.T) {
	// sample/ predates the Tailwind chassis and leans on its own style.css.
	root := t.TempDir()
	write(t, filepath.Join(root, cssScript), "#!/bin/sh\n")
	if err := os.MkdirAll(filepath.Join(root, "talk", ".demoit"), 0o755); err != nil {
		t.Fatalf("unable to create the talk: %v", err)
	}
	chdir(t, root)

	builder, reason := newStylesheetBuilder("talk")
	if builder != nil {
		t.Fatalf("got a builder, want none for a talk with no tailwind.src.css")
	}
	if !strings.Contains(reason, "tailwind.src.css") {
		t.Errorf("got reason %q, want it to name the missing build entry", reason)
	}
}

func write(t *testing.T, path, content string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("unable to create %s: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("unable to write %s: %v", path, err)
	}
}

func TestOnlyASourceFileTriggersARebuild(t *testing.T) {
	t.Parallel()

	for path, want := range map[string]bool{
		"sdd-talk/demoit.md":              true,
		"sdd-talk/talk.yml":               true,
		"sdd-talk/.demoit/style.css":      true,
		"sdd-talk/.demoit/layouts/x.html": true,
		"sdd-talk/.demoit/images/a.png":   false,
		"sdd-talk/.demoit/js/demoit.js":   false,
		// The builder writes this one inside the folder being watched, so
		// treating it as a source would have every build trigger the next.
		"sdd-talk/.demoit/" + stylesheetOutput: false,
	} {
		if got := triggersRebuild(path); got != want {
			t.Errorf("triggersRebuild(%q) = %v, want %v", path, got, want)
		}
	}
}
