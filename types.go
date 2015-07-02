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
	"errors"
	"fmt"
	"strings"
)

type MissingURLParameterError string

func (e MissingURLParameterError) Error() string {
	return fmt.Sprintf("apidoc: missing documentation for URL param: %s", string(e))
}

var (
	ErrMissingMethod = errors.New("apidoc: missing HTTP verb")
	ErrMissingURL    = errors.New("apidoc: missing URL")
)

// An Endpoint represents the pertinent documentatopn for a single HTTP API endpoint.
type Endpoint struct {

	// Description is a human-readable description of the parameter and it's
	// functionality
	Description string

	// Method is the HTTP request verb: e.g. GET, PUT, POST, DELETE
	Method string

	// URLTemplate is the URL structure of the endpoint, showing any URL params
	// with colons, e.g. "/foobar/v1/hello/:firstName/:lastName"
	URLTemplate string

	// URLParams are the set of parameters that are specified in the URL of a
	// request.
	URLParams []Parameter

	// DataParams are the set of parameters that are specified in the body of a
	// request.
	DataParams []Parameter

	// SuccessResponse is a description of the response code and response body
	// that a client can expect on a successful call to the Endpoint
	SuccessResponse Response

	// ErrorResponses are descriptions of the response codes and response bodies
	// that a client can expect on a failed call to the Endpoint
	ErrorResponses []Response

	// Examples are meant to to illustrate a syntactically-correct example call
	// to the given endpoint. A common use case is showing a curl command.
	Examples []string

	// Notes is a description of any important behaviors, side-effects, or other
	// pertinent details of the endpoint
	Notes string
}

// Validate ensure that all of the required fields are valid for an Endpoint.
// Currently that simply means: the HTTP method and URL are specified, and
// any params referenced in the URL have corresponding Parameter instances.
func (e Endpoint) Validate() error {
	if e.Method == "" {
		return ErrMissingMethod
	}

	if e.URLTemplate == "" {
		return ErrMissingURL
	}

	for _, split := range strings.Split(e.URLTemplate, "/") {
		if strings.HasPrefix(split, ":") && !contains(e.URLParams, split[1:]) {
			return MissingURLParameterError(split[1:])
		}
	}
	return nil
}

func contains(ps []Parameter, name string) bool {
	for _, p := range ps {
		if p.Name == name {
			return true
		}
	}
	return false
}

// A Response represents a type of HTTP response from an Endpoint.
type Response struct {

	// Code is the HTTP response code of the Response
	Code int

	// ExampleContent shows a representative response body
	Content string
}

// A Parameter represents either a URL parameter or a request body parameter
// for an Endpoint
type Parameter struct {

	// Name is the identifier for the parameter.  It should correspond with a
	// segment of the URLTemplate in an Endpoint (for GET-type requests) or an
	// attribute of the request body (for POST|PUT-type requests)
	Name string

	// Required indicates whether or not the parameter is required to be
	// specified
	Required bool

	// Type indicates the type of the parameter. e.g. string, numeric, etc.
	Type string

	// Description is a human-readable description of the parameter and it's
	// functionality
	Description string
}
