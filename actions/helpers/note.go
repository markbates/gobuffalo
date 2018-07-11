package helpers

import (
	"fmt"
	"html/template"

	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/tags"
	"github.com/pkg/errors"
	"github.com/shurcooL/github_flavored_markdown"
)

// Note is a template block helper to insert a remark.
func Note(help plush.HelperContext) (template.HTML, error) {
	if !help.HasBlock() {
		return "", nil
	}
	s, err := help.Block()
	if err != nil {
		return "", errors.WithStack(err)
	}

	b := github_flavored_markdown.Markdown([]byte(s))

	// Generate info div
	// Strip the first <p> from Markdown result to insert the icon.
	t := tags.New("div", tags.Options{
		"class": "info",
		"body":  fmt.Sprintf(note, b[3:]),
	})

	return t.HTML(), nil
}

// Warning is a template block helper to insert a warning.
func Warning(help plush.HelperContext) (template.HTML, error) {
	if !help.HasBlock() {
		return "", nil
	}
	s, err := help.Block()
	if err != nil {
		return "", errors.WithStack(err)
	}

	b := github_flavored_markdown.Markdown([]byte(s))

	// Generate info div
	// Strip the first <p> from Markdown result to insert the icon.
	t := tags.New("div", tags.Options{
		"class": "warning",
		"body":  fmt.Sprintf(warning, b[3:]),
	})

	return t.HTML(), nil
}

const note = `<p><i class="fa fa-info-circle" aria-hidden="true"></i>%s`
const warning = `<i class="fa fa-exclamation-triangle" aria-hidden="true"></i>%s`
