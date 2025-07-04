package scaffolder

import (
	"bytes"
	"embed"
	"path/filepath"
	"strings"
	"text/template"

	u "github.com/Rick-Phoenix/goutils"
	"github.com/labstack/gommon/log"
)

//go:embed templates/*
var templateFS embed.FS

func Scaffold() {
	tmpl, err := template.New("scaffolder").Funcs(funcMap).ParseFS(templateFS, "templates/*")
	if err != nil {
		log.Fatalf("Could not initialize the template instance: %s", err.Error())
	}

	for _, tem := range tmpl.Templates() {
		var buffer bytes.Buffer
		err := tem.Execute(&buffer, "")
		if err != nil {
			log.Fatal(err)
		}

		ext := filepath.Ext(tem.Name())
		outPath := filepath.Join("gen", strings.TrimSuffix(tem.Name(), ext))
		err = u.WriteFile(buffer, outPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

var funcMap = template.FuncMap{}
