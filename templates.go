package u

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

func FormatFile(path string) error {
	fmtCmd := "gofmt"
	_, notFoundErr := exec.LookPath("gofumpt")
	if notFoundErr == nil {
		fmtCmd = "gofumpt"
	}
	cmd := exec.Command(fmtCmd, "-w", path)
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Could not format the file at %q:\n%s\n", path, err.Error())
	}

	return nil
}

func RunGoImports(path string) error {
	importCmd := exec.Command("goimports", "-w", path)
	importCmd.Stderr = os.Stderr
	importErr := importCmd.Run()
	if importErr != nil {
		fmt.Printf("An error occurred while calling goimports for the file %q:\n%s\n", path, importErr.Error())
	}

	return nil
}

func FormatAndImports(path string, fatal bool) error {
	var error error
	err := FormatFile(path)
	if err != nil {
		err = fmt.Errorf("Error while trying to format the file at %q:\n%w\n", path, err)
		if fatal {
			return err
		} else {
			fmt.Print(err)
		}
	}

	impErr := RunGoImports(path)
	if impErr != nil {
		impErr = fmt.Errorf("Error while trying to run goimports for the file at %q:\n%w\n", path, impErr)
		if fatal {
			return impErr
		} else {
			fmt.Print(impErr)
		}
	}

	return error
}

func ExecTemplate(tmpl *template.Template, templateName string, outputPath string, context any) error {
	var buffer bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buffer, templateName, context); err != nil {
		return fmt.Errorf("Failed to execute template %q:\n%w", templateName, err)
	}

	if err := WriteFile(buffer, outputPath); err != nil {
		return err
	}

	return nil
}

func ExecTemplateAndFormat(tmpl *template.Template, templateName string, outputPath string, context any) error {
	err := ExecTemplate(tmpl, templateName, outputPath, context)
	if err != nil {
		return err
	}

	err = FormatAndImports(outputPath, false)
	return err
}
