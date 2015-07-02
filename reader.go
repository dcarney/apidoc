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
	"fmt"
	"go/ast"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	KWDescription     = "Description"
	KWNotes           = "Notes"
	KWSuccessResponse = "Success Response"
	KWErrorResponse   = "Error Response"
	KWExample         = "Example"
	KWParameter       = "Parameter"
	KWMethod          = "Method"
	KWNone            = "(none)"
)

var (
	apidocMarker    = `(apidoc)\(([^)]+)\):?`                           // apidoc(name), name of at least 1 char
	apidocMarkerRx  = regexp.MustCompile(`^[ \t]*` + apidocMarker)      // the marker at text start
	apidocCommentRx = regexp.MustCompile(`^/[/*][ \t]*` + apidocMarker) // the marker at comment start

	// Example: GET /some/path/:foo
	httpVerbRx = regexp.MustCompile(`(GET|PUT|POST|DELETE|HEAD|OPTIONS|TRACE|CONNECT|PATCH)\s+(/.*)`)

	// Parameter docs follow the pattern:   name [, "required"] [, type]
	// Examples:
	//    foobar
	//    foobar, string
	//    foobar, required, string
	//    foobar, required
	//    foobar, required, array of strings
	parameterRx = regexp.MustCompile(`^([\w-]+)(?:\s*,\s*(required))?(?:\s*,\s*([\w\s]+))?$`)
)

// TODO: change this to a regexp?
func startsWithKeyword(str string) string {

	switch {
	case strings.HasPrefix(str, KWDescription):
		return KWDescription
	case strings.HasPrefix(str, KWSuccessResponse):
		return KWSuccessResponse
	case strings.HasPrefix(str, KWErrorResponse):
		return KWErrorResponse
	case strings.HasPrefix(str, KWExample):
		return KWExample
	case strings.HasPrefix(str, KWParameter):
		return KWParameter
	case strings.HasPrefix(str, KWNotes):
		return KWNotes
	case httpVerbRx.MatchString(str):
		return KWMethod
	}
	return KWNone
}

// stripKeyword removes the given keyword from the first line, as well as any
// leading whitespace-only lines
func stripKeyword(kw string, lines []string) []string {
	lines[0] = strings.TrimSpace(strings.TrimPrefix(lines[0], kw))
	i := -1
	for j, line := range lines {
		if strings.TrimSpace(line) != "" {
			i = j
			break
		}
	}

	if i > 0 {
		return lines[i:]
	}
	return lines
}

// parseKeyword populates Endpoint fields, based on the content in the
// supplied comment lines.
func parseKeyword(e *Endpoint, kw string, lines []string) error {

	switch kw {
	case KWMethod:
		matches := httpVerbRx.FindStringSubmatch(lines[0])
		if len(matches) > 0 {
			e.Method = matches[1]
			e.URLTemplate = matches[2]
		}
	case KWDescription:
		lines = stripKeyword(KWDescription, lines)
		e.Description = strings.Join(lines, "\n")
	case KWParameter:
		lines = stripKeyword(KWParameter, lines)
		matches := parameterRx.FindStringSubmatch(lines[0])
		if len(matches) > 0 {
			p := Parameter{
				Name:        matches[1],
				Required:    matches[2] == "required",
				Type:        matches[3],
				Description: strings.Join(lines[1:], " "),
			}

			e.URLParams = append(e.URLParams, p)
		}

	case KWSuccessResponse:
		lines = stripKeyword(KWSuccessResponse, lines)
		code, err := strconv.Atoi(lines[0])
		if err != nil {
			return err
		}
		e.SuccessResponse = Response{
			Code:    code,
			Content: strings.Join(lines[1:], "\n"),
		}
	case KWErrorResponse:
		lines = stripKeyword(KWErrorResponse, lines)
		code, err := strconv.Atoi(lines[0])
		if err != nil {
			return err
		}
		er := Response{
			Code:    code,
			Content: strings.Join(lines[1:], "\n"),
		}
		e.ErrorResponses = append(e.ErrorResponses, er)
	case KWExample:
		lines = stripKeyword(KWExample, lines)
		e.Examples = append(e.Examples, strings.Join(lines, "\n"))
	case KWNotes:
		lines = stripKeyword(KWNotes, lines)
		e.Notes = strings.Join(lines, "\n")
	default:
		return fmt.Errorf("Unknown keyword: %s", kw)
	}

	return nil
}

// parseEndpoint takes an apidoc body (which consists of one or more
// newline-separated lines) and parses the various keyword sections, populating
// an Endpoint.  The body for each keyword extends until the next keyword,
// or until the end of the body.
func parseEndpoint(body string) (*Endpoint, error) {
	lines := strings.Split(body, "\n")

	e := &Endpoint{}
	i := -1
	lastKw := KWNone
	for j, line := range lines {
		kw := startsWithKeyword(line)
		if kw != KWNone {
			if i >= 0 {
				if err := parseKeyword(e, lastKw, lines[i:j]); err != nil {
					return e, err
				}
			}
			i = j
			lastKw = kw
		}
	}
	if i >= 0 {
		if err := parseKeyword(e, lastKw, lines[i:]); err != nil {
			return e, err
		}
	}

	return e, nil
}

// A reader read a series of CommentGroups, looking for, and attempting to parse
// apidoc text blocks.
type reader struct {
	strict    bool
	endpoints []*Endpoint
}

// readDocs extracts apidoc from comments.  An apidoc must start at the
// beginning of a comment with "apidoc(name):" and is followed by the lines
// that make up the body.  The apidoc ends at the end of the comment group or
// at the start of another apidoc in the same comment group, whichever comes
// first.
func (r *reader) readDocs(comments []*ast.CommentGroup) error {
	for _, group := range comments {
		i := -1 // comment index of most recent note start, valid if >= 0
		list := group.List
		for j, c := range list {
			if apidocCommentRx.MatchString(c.Text) {
				if i >= 0 {
					if err := r.readDoc(list[i:j]); err != nil {
						return err
					}
				}
				i = j
			}
		}
		if i >= 0 {
			if err := r.readDoc(list[i:]); err != nil {
				return err
			}
		}
	}
	return nil
}

// readDoc collects a single api doc from a sequence of comments.
func (r *reader) readDoc(list []*ast.Comment) error {
	text := (&ast.CommentGroup{List: list}).Text()
	if m := apidocMarkerRx.FindStringSubmatchIndex(text); m != nil {
		// The doc body starts after the marker.
		body := text[m[1]:]
		if body != "" {
			e, err := parseEndpoint(body)
			if err != nil {
				return err
			}
			if err := e.Validate(); err != nil {
				if r.strict {
					log.Fatalf("validation error: %s\n", err.Error())
				} else {
					log.Printf("validation error: %s\n", err.Error())
				}
			}
			r.endpoints = append(r.endpoints, e)
		}
	}
	return nil
}
