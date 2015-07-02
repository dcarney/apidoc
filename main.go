package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type renderFunc func(*Endpoint, io.Writer) error

func processFile(inputPath, outputPath string, render renderFunc) {
	var err error
	fmt.Printf("process file: %s, %s\n", inputPath, outputPath)
	endpoints := loadFile(inputPath)
	if len(endpoints) == 0 {
		return
	}

	out, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("could not open output file: %s", err)
	}
	defer out.Close()

	log.Printf("rendering template to: %s\n", outputPath)
	for _, endpoint := range endpoints {
		if err = render(endpoint, out); err != nil {
			log.Fatalf("could not generate apidoc: %s", err)
		}
	}
}

func loadFile(inputPath string) []*Endpoint {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, inputPath, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("error parsing file: %s", err)
	}

	r := reader{
		strict: opts.strict,
	}
	if err := r.readDocs(f.Comments); err != nil {
		log.Printf("error reading apidoc in src file %s: %s\n", inputPath, err.Error())
	}

	if len(r.endpoints) > 0 {
		log.Printf("found %d apidoc(s) in src file: '%s'\n", len(r.endpoints), inputPath)
	}
	return r.endpoints
}

func deriveOutputPath(inputPath, extension string) string {
	dir, fName := filepath.Split(strings.TrimSuffix(inputPath, ".go"))
	gopkg := os.Getenv("GOPACKAGE")
	if gopkg != "" {
		return filepath.Join(dir, fmt.Sprintf("%s_%s_apidoc.%s", gopkg, fName, extension))
	}

	return filepath.Join(dir, fmt.Sprintf("%s_apidoc.%s", fName, extension))
}

var opts struct {
	strict bool
	output string
	format int
}

func main() {

	var format string
	flag.StringVar(&opts.output, "out", "", "Name of the output file to use. If not specified, the output file name will be based on the package and input file name.")
	flag.BoolVar(&opts.strict, "strict", false, "Enables validation checks on each apidoc comment block. When strict is true, any validation error causes the process to exit.")
	flag.StringVar(&format, "format", "markdown", "Specifies the format to render the docs in [markdown|html]. Defaults to markdown.")
	flag.Parse()

	log.SetFlags(0)
	log.SetPrefix("apidoc: ")

	var render renderFunc
	var ext string
	switch format {
	case "html":
		render = RenderHtml
		ext = "html"
	case "markdown":
		render = RenderMarkdown
		ext = "md"
	default:
		log.Fatalf("invalid format '%s'. Form can be: [markdown|html]", format)
	}

	if opts.strict {
		log.Println("strict mode")
	}

	for _, path := range flag.Args() {
		if !strings.HasSuffix(path, ".go") {
			panic(fmt.Errorf("input file %s doesn't have .go extension", path))
		}

		if opts.output == "" {
			processFile(path, deriveOutputPath(path, ext), render)
		} else {
			processFile(path, opts.output, render)
		}
	}
}
