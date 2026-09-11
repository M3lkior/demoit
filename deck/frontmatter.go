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

package deck

import "gopkg.in/yaml.v3"

// knownKeys are the frontmatter keys the engine acts on. Every other key is
// handed to the layout through Slide.Meta.
var knownKeys = []string{"layout", "title", "source", "sources", "class", "height", "speakernotes"}

// MaxSources is how many entries a slide's sources: list may carry. The cap is
// a readability one, not a technical one: the list renders on a single line at
// the foot of the slide, and past five entries that line wraps into the slide
// itself. Going over is reported as a slide error rather than silently
// truncated, so an author sees which slide to trim instead of losing a
// reference without being told.
const MaxSources = 5

// frontmatter holds the frontmatter keys the engine acts on.
type frontmatter struct {
	Layout string `yaml:"layout"`
	Title  string `yaml:"title"`
	// Source is the original single-reference key, kept for decks that use it.
	Source string `yaml:"source"`
	// Sources is the list form, rendered as a bullet list at the foot of the
	// slide. A slide may carry both; the layout renders both.
	Sources      []string `yaml:"sources"`
	Class        string   `yaml:"class"`
	Height       string   `yaml:"height"`
	Speakernotes string   `yaml:"speakernotes"`
}

// parseFrontmatter reads a slide's YAML block into the keys the engine knows
// about, plus a map of everything else.
func parseFrontmatter(src []byte) (frontmatter, map[string]any, error) {
	var known frontmatter
	if len(src) == 0 {
		return known, nil, nil
	}

	if err := yaml.Unmarshal(src, &known); err != nil {
		return known, nil, err
	}

	rest := map[string]any{}
	if err := yaml.Unmarshal(src, &rest); err != nil {
		return known, nil, err
	}
	for _, key := range knownKeys {
		delete(rest, key)
	}

	return known, rest, nil
}
