package scaffolder

import (
	"bytes"
	"embed"
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	u "github.com/Rick-Phoenix/goutils"
)

//go:embed templates
var templateFS embed.FS

func ScaffoldMoonRepo(rootDir string, tmplData any) error {
	tmpl, err := parseTemplates("moon/*")
	if err != nil {
		return err
	}

	err = writeTemplates(tmpl, rootDir, tmplData)
	if err != nil {
		return err
	}

	u.LogSuccess(fmt.Sprintf("Generated moonrepo files in %s", rootDir))

	return nil
}

func ScaffoldGo(rootDir string, tmplData any) error {
	tmpl, err := parseTemplates(".gitignore", ".pre-commit-config.yaml", "Taskfile.yaml", "main_test.go")
	if err != nil {
		return err
	}

	err = writeTemplates(tmpl, rootDir, tmplData)
	if err != nil {
		return err
	}

	u.LogSuccess(fmt.Sprintf("Scaffolded go project in %s", rootDir))

	return nil
}

func ScaffoldSvelte(rootDir string, tmplData any) error {
	tmpl, err := parseTemplates("svelte/*/*/*", "svelte/*")
	if err != nil {
		return err
	}

	err = writeTemplates(tmpl, rootDir, tmplData)
	if err != nil {
		return err
	}

	u.LogSuccess(fmt.Sprintf("Generated svelte project in %s", rootDir))

	return nil
}

func parseTemplates(pattern ...string) (*template.Template, error) {
	fullPatterns := make([]string, len(pattern))
	for i, pa := range pattern {
		fullPatterns[i] = filepath.Join("templates", pa+".tmpl")
	}
	tmpl, err := template.New("scaffolder").Funcs(funcMap).ParseFS(templateFS, fullPatterns...)
	if err != nil {
		return nil, fmt.Errorf("Could not initialize the template instance: %w", err)
	}

	return tmpl, nil
}

func debugTemplates(tmpl *template.Template) {
	u.LogDebug(fmt.Sprintf("Templates in %s", tmpl.Name()))
	for _, tem := range tmpl.Templates() {
		fmt.Printf("%+v\n", tem.Name())
	}
}

func writeTemplates(tmpl *template.Template, rootDir string, tmplData any, ignores ...string) error {
	for _, tem := range tmpl.Templates() {
		skip := slices.Contains(ignores, tem.Name())
		if strings.HasSuffix(tem.Name(), ".tmpl") {
			skip = true
		}
		if skip {
			continue
		}
		var buffer bytes.Buffer
		err := tmpl.ExecuteTemplate(&buffer, tem.Name(), tmplData)
		if err != nil {
			return err
		}

		outputPath := filepath.Join(rootDir, tem.Name())

		confirm := u.PromptIfFileExists(outputPath)

		if confirm {
			err = u.WriteFile(buffer, outputPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

var funcMap = template.FuncMap{}
