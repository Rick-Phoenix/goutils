package scaffolder

import (
	"bytes"
	"embed"
	"path/filepath"
	"text/template"

	u "github.com/Rick-Phoenix/goutils"
	"github.com/labstack/gommon/log"
)

//go:embed templates/*
var templateFS embed.FS

func Scaffold(templates []string, tmplData any) {
	tmpl, err := template.New("scaffolder").Funcs(funcMap).ParseFS(templateFS, "templates/*.tmpl", "templates/**/*.tmpl")
	if err != nil {
		log.Fatalf("Could not initialize the template instance: %s", err.Error())
	}

	// for _, tem := range tmpl.Templates() {
	// 	fmt.Printf("DEBUG: %+v\n", tem.Name())
	// }

	for _, tem := range templates {
		var buffer bytes.Buffer
		err := tmpl.ExecuteTemplate(&buffer, tem, tmplData)
		if err != nil {
			log.Fatal(err)
		}

		confirm := u.PromptIfFileExists(tem)

		if confirm {
			err = u.WriteFile(buffer, tem)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func ScaffoldSvelte(rootDir string, tmplData any) {
	tmpl, err := template.New("scaffolder").Funcs(funcMap).ParseFS(templateFS, "templates/svelte/**/*.tmpl")
	if err != nil {
		log.Fatalf("Could not initialize the template instance: %s", err.Error())
	}

	// for _, tem := range tmpl.Templates() {
	// 	fmt.Printf("DEBUG: %+v\n", tem.Name())
	// }

	for _, tem := range tmpl.Templates() {
		var buffer bytes.Buffer
		err := tmpl.ExecuteTemplate(&buffer, tem.Name(), tmplData)
		if err != nil {
			log.Fatal(err)
		}

		confirm := u.PromptIfFileExists(tem.Name())

		if confirm {
			err = u.WriteFile(buffer, filepath.Join(rootDir, tem.Name()))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

var funcMap = template.FuncMap{}
