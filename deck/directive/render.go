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

package directive

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// spec says how a directive name maps to HTML.
type spec struct {
	// tag is the custom element the directive renders as.
	tag string
	// container reports whether the element wraps children.
	container bool
	// attributes renders the element's attributes, leading space included.
	attributes func(attrs map[string]string) (string, error)
	// content is the text written between the tags of a non-container. Only
	// ::stars needs it: buttons.github.io reads the label of its <a>, so the
	// element cannot be written empty the way every custom element here is.
	content string
}

// specs is the catalogue of directives, keyed by the name written in Markdown.
var specs = map[string]spec{
	"split":        {tag: "split-view", container: true, attributes: splitAttributes},
	"window":       {tag: "fake-window", container: true, attributes: copyAttributes("title")},
	"speakernotes": {tag: "speaker-notes", container: true, attributes: copyAttributes()},
	"term":         {tag: "web-term", attributes: copyAttributes("path")},
	"browser":      {tag: "web-browser", attributes: copyAttributes("src")},
	"vscode":       {tag: "vs-code", attributes: copyAttributes("path")},
	"code":         {tag: "source-code", attributes: codeAttributes},
	"grid":         {tag: "div", container: true, attributes: gridAttributes},
	"col":          {tag: "div", container: true, attributes: gridAttributes},
	"stars":        {tag: "a", attributes: starsAttributes, content: "Star"},
}

// nodeRenderer renders directive nodes as demoit custom elements. It carries
// the parse context so that an attribute problem is reported as a slide-level
// error instead of failing the whole conversion.
type nodeRenderer struct {
	ctx parser.Context
}

// RegisterFuncs implements renderer.NodeRenderer.
func (r *nodeRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(Kind, r.render)
}

// render writes the element a directive maps to.
func (r *nodeRenderer) render(w util.BufWriter, _ []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	directiveNode, ok := node.(*Node)
	if !ok {
		return ast.WalkContinue, nil
	}

	element, known := specs[directiveNode.Name]
	if !known {
		// Unknown names are reported by the transformer; render nothing.
		return ast.WalkSkipChildren, nil
	}

	if !entering {
		if element.container {
			_, _ = w.WriteString("</" + element.tag + ">\n")
		}

		return ast.WalkContinue, nil
	}

	attributes, err := element.attributes(directiveNode.Attrs)
	if err != nil {
		addError(r.ctx, directiveNode.Line, "%s", err.Error())

		return ast.WalkSkipChildren, nil
	}

	_, _ = w.WriteString("<" + element.tag + attributes + ">")
	if !element.container {
		_, _ = w.WriteString(element.content + "</" + element.tag + ">\n")
	}

	return ast.WalkContinue, nil
}

// copyAttributes copies the named attributes through, in order, skipping the
// ones that are absent or empty. Values are HTML-escaped: %q would quote them
// the Go way, which leaves a " in a value free to end the attribute and turn
// whatever follows into markup of its own.
func copyAttributes(keys ...string) func(map[string]string) (string, error) {
	return func(attrs map[string]string) (string, error) {
		rendered := ""
		for _, key := range keys {
			if value := attrs[key]; value != "" {
				rendered += fmt.Sprintf(` %s="%s"`, key, html.EscapeString(value))
			}
		}

		return rendered, nil
	}
}

// splitAttributes renders the class of a <split-view>, built from its height.
func splitAttributes(attrs map[string]string) (string, error) {
	if height := attrs["height"]; height != "" {
		return fmt.Sprintf(` class="%s"`, html.EscapeString(stageHeightClass(height))), nil
	}

	return "", nil
}

// stageHeightClass is the height utility a height name maps to. The name
// travels verbatim -- h-stage-xlarge, not h-stage-xl -- because a deck writes
// `height: xlarge` and deck/layouts/split.html interpolates that value as-is;
// translating here would mean maintaining a lookup table for nothing.
//
// The tokens behind these utilities are a share of the slide's main area rather
// than an absolute length, so a pane cannot be taller than the stage holding
// it. It is shared with the transformer, which appends the same class to a
// weighted grid.
func stageHeightClass(height string) string {
	return "h-stage-" + height
}

// gridAttributes renders the class of a grid or column node, whether the
// transformer built it from a split{cols=} or the author wrote :::grid /
// :::col by hand. The value is escaped for the same reason as its siblings
// above: a " in a hand-written class ends the attribute early and turns the
// rest of the line into markup.
func gridAttributes(attrs map[string]string) (string, error) {
	class := strings.TrimSpace(attrs["class"])

	// `reveal` marks the whole block as one step of a progressive reveal. It
	// is a flag, written bare, and rendered as a class rather than as its own
	// attribute so that the stylesheet and demoit.js have a single thing to
	// look for, whether the author marked a directive or wrote the class on a
	// plain HTML tag themselves.
	//
	// Only a column or a grid takes it here, and that is the point of putting
	// it on this function alone: those are the wrappers an author reaches for
	// when a step is more than one element. Anything else is a tag the author
	// writes, and a tag takes class="reveal" directly.
	if _, marked := attrs["reveal"]; marked {
		class = strings.TrimSpace(class + " reveal")
	}

	if class != "" {
		return fmt.Sprintf(` class="%s"`, html.EscapeString(class)), nil
	}

	return "", nil
}

// starsAttributes renders a GitHub star button for `::stars{repo=owner/repo}`.
//
// The repo is the whole input, and it is checked rather than escaped into a
// URL and hoped for: it is interpolated into an href, so anything but the
// owner/repo shape it claims to be is a slide error. That keeps a typo on the
// slide, where the author sees it, instead of in a link that goes somewhere
// unintended.
func starsAttributes(attrs map[string]string) (string, error) {
	repo := strings.TrimSpace(attrs["repo"])
	if repo == "" {
		return "", errors.New("stars: repo is required, as repo=owner/name")
	}
	if !isRepoPath(repo) {
		return "", fmt.Errorf("stars: %q is not an owner/name repository path", repo)
	}

	return fmt.Sprintf(` class="github-button" href="https://github.com/%s" data-size="large" data-show-count="true" aria-label="Star %s on GitHub"`,
		html.EscapeString(repo), html.EscapeString(repo)), nil
}

// isRepoPath reports whether path is a bare `owner/name`, with the characters
// GitHub allows in each half and nothing else -- no scheme, no host, no extra
// segment, no traversal.
func isRepoPath(path string) bool {
	owner, name, found := strings.Cut(path, "/")
	if !found || owner == "" || name == "" {
		return false
	}

	for _, half := range []string{owner, name} {
		for _, c := range half {
			switch {
			case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z', c >= '0' && c <= '9':
			case c == '-', c == '_', c == '.':
			default:
				return false
			}
		}
	}

	return true
}
