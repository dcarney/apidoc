package boilerdoc

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
	// request
	DataParams []Parameter

	// SuccessResponse is a description of the response code and response body
	// that a client can expect on a successful call to the Endpoint
	SuccessResponse Response

	// ErrorResponses are descriptions of the response codes and response bodies
	// that a client can expect on a failed call to the Endpoint
	ErrorResponses []Response

	// ExampleCalls are examples meant to to illustrate a syntactically-correct
	// example call to the given endpoint
	ExampleCalls []string

	// Notes is a description of any important behaviors, side-effects, or other
	// pertinent details of the endpoint
	Notes string
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
