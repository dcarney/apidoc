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

const (
	htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>apidoc</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css">
  </head>
  <body>
		<div class="container">
			<h3> {{ .Method }} [{{ .URLTemplate }}] </h3>

			<p>{{ .Description }}</p>

			{{ if .Notes }}
			<p><em>**NOTE:**</em> {{ .Notes }}</p>
			{{end}}

			{{ if .URLParams }}
			<h4>Parameters</h4>
			<ul>
				{{ range $param := .URLParams }}
				  <li>{{ $param.Name }} ({{ if $param.Required }}required {{ end }}{{ if $param.Type }}{{ $param.Type }}{{ end }}) : {{ $param.Description }}</li>
				{{ end }}
			</ul>
			{{ end }}

			{{ if .DataParams }}
			<h4>Request Body Parameters</h4>
      <ul>
				{{ range $param := .DataParams }}
				  <li>{{ $param.Name }} ({{ if $param.Required }}required {{ end }}{{ if $param.Type }}{{ $param.Type }}{{ end }}) : {{ $param.Description }}</li>
				{{ end }}
      </ul>
			{{ end }}

			<h4>Example success response</h4>
			<code>{{ .SuccessResponse.Code }}</code>:<span>{{ statusText .SuccessResponse.Code }}</span>
			<pre>{{ .SuccessResponse.Content }}</pre>

			{{ if .ErrorResponses }}
			<h4>Example error responses</h4>
				{{ range $resp := .ErrorResponses }}
				<code>{{ $resp.Code }}</code>:<span>{{ statusText $resp.Code }}</span>
				<pre>{{ $resp.Content }}</pre>
				{{ end }}
			{{ end }}

			{{ if .Examples }}
			<h4>Examples</h4>
				{{ range $call := .Examples }}
					<pre>{{ $call }}</pre>
				{{ end }}
			{{ end }}

    </div>
  </body>
</html>
`
)
