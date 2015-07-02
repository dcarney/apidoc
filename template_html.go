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
