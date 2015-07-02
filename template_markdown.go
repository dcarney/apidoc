package main

const (

	// the template for Markdown output.  Unfortunately, there's no escaping a backtick(`) inside backticks in Go,
	// so we have to use concatenation to get backtick literals:
	// e.g. `some string` + "`" + `another string` + "`"
	markdownTemplate = `
### {{ .Method }} [{{ .URLTemplate }}]

{{ .Description }}

{{ if .Notes }}
**NOTE:** {{ .Notes }}
{{end}}

{{ if .URLParams }}
#### Parameters
  {{ range $param := .URLParams }}
  * {{ $param.Name }} ({{ if $param.Required }}required {{ end }}{{ if $param.Type }}{{ $param.Type }}{{ end }}) : {{ $param.Description }}
  {{ end }}
{{ end }}

{{ if .DataParams }}
#### Request Body Parameters
  {{ range $param := .DataParams }}
  * {{ $param.Name }} ({{ if $param.Required }}required {{ end }}{{ if $param.Type }}{{ $param.Type }}{{ end }}) : {{ $param.Description }}
  {{ end }}
{{ end }}

#### Example success response
` + "`" + `{{ .SuccessResponse.Code }}` + "`" + `: {{ statusText .SuccessResponse.Code }}

    {{ .SuccessResponse.Content }}

{{ if .ErrorResponses }}
#### Example error responses
  {{ range $resp := .ErrorResponses }}
  ` + "`" + `{{ $resp.Code }}` + "`" + `: {{ statusText $resp.Code }}

    {{ $resp.Content }}
  {{ end }}
{{ end }}

{{ if .Examples }}
#### Examples
  {{ range $call := .Examples }}
    {{ $call }}
  {{ end }}
{{ end }}
`
)
