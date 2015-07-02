// The MIT License (MIT)
//
// Copyright (c) 2015 Dylan Carney
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
