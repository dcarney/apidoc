package boilerdoc

import (
	"io"
	"net/http"
	"text/template"
)

// Render writes a Markdown representation of the specified Endpoint to an
// io.Writer using the supplied templateFile.
func Render(e Endpoint, out io.Writer, templateFile string) error {
	funcs := template.FuncMap{
		"statusText": http.StatusText,
	}

	t := template.Must(template.New("endpoint.template").Funcs(funcs).ParseFiles(templateFile))
	return t.Execute(out, e)
}
