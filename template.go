package main

import (
	"io"
	"net/http"
	"text/template"
)

// RenderMarkdown writes a Markdown representation of the specified Endpoint to
// an io.Writer
func RenderMarkdown(e *Endpoint, out io.Writer) error {
	fm := template.FuncMap{
		"statusText": http.StatusText,
	}
	t := template.Must(template.New("markdown").Funcs(fm).Parse(markdownTemplate))
	return t.Execute(out, e)
}

// RenderHtml writes a Markdown representation of the specified Endpoint to
// an io.Writer
func RenderHtml(e *Endpoint, out io.Writer) error {
	fm := template.FuncMap{
		"statusText": http.StatusText,
	}
	t := template.Must(template.New("html").Funcs(fm).Parse(htmlTemplate))
	return t.Execute(out, e)
}
