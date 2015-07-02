//go:generate apidoc -strict=true $GOFILE
package main

import "net/http"

// apidoc(foobar)
//
// GET /someapi/v1/:foo/:bar
//
// Description
// This is a description of the endpoint and what it does.  There
// are multiple sentences and lines that are included.
//
// Parameter foo, required, string
// This is a description of foo
//
// Parameter bar, required, string
// And this is a description of bar
//
// Success Response 200
// 		{ "Message": "this shows a 200 response!" }
//
// Error Response 400
//    { "Message": "that's a bad request" }
//
// Error Response 401
//    { "Message": "is your auth correct?" }
//
// Example
// 		curl -v -k -H "X-SomeCustomHeader: 12345" https://host/someapi/v1/foo/bar
//
// Foobar is an http.Handler that does things
func Foobar(w http.ResponseWriter, r *http.Request) {
}

// FizzBuzz is an http.Handler that handles multiple API requests.  This is the normal godoc.
// There can be multiple apidoc() sections, as shown below:
//
// apidoc(fizz)
//
// POST /someapi/v1/fizz/:foo
//
// Description
// This is a description of the endpoint and what it does.  There
// are multiple sentences and lines that are included.
//
// This is still description.
//
// Parameter foo, string
// Foo has a type, but is not required
//
// Parameter bar, required
// Bar is required, but it doesn't have a type
//
// Success Response 200
// 		{ "Message": "this shows a 200 response!" }
//
// Error Response 400
//    { "Message": "that's a bad request" }
//
// Example
// 		curl -v -H https://host/someapi/v1/fizz/buzz -d '
//    {
//      "attr": "foo",
//      "attrB": "bar",
//      "attrC": 89.45
//		}'
//
// apidocs(buzz)
//
// GET /someapi/v2/something/:fizz/:buzz
//
// Description
// This is a description of the endpoint and what it does.  There
// are multiple sentences and lines that are included.
//
// A section
//
// This is still the description, since no other keyword has been found yet.
// The section header above is just normal godoc behavior.
//
//
// Another section
//
// Even more description. This goes on and on and on until another keyword is
// found, or until the end of the comment group.
//
// Parameter
// fizz, required
// This is a description
// of fizz
//
// that takes up
// many
//
// lines
//
// Parameter buzz, required, string
// And this is a description of buzz
//
// Success Response 200
// 		{ "Message": "this shows a 200 response!" }
//
// Error Response 400
//    { "Message": "that's a bad request" }
//
//
func FizzBuzz(w http.ResponseWriter, r *http.Request) {}
