package u

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func WriteFile(buffer bytes.Buffer, path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("Could not create the missing directories to write the file %q:\n%w\n", path, err)
	}

	if err := os.WriteFile(path, buffer.Bytes(), 0644); err != nil {
		return fmt.Errorf("Could not write the file %q:\n%w\n", path, err)
	}

	return nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func PromptIfFileExists(path string) bool {
	exists := FileExists(path)
	if exists {
		fmt.Printf("File '%s' already exists. Overwrite? (y/N): ", path)
		reader := bufio.NewScanner(os.Stdin)
		reader.Scan()
		response := strings.ToLower(strings.TrimSpace(reader.Text()))

		if response != "y" {
			fmt.Printf("Skipping file '%s'.\n", path)
			return false
		}
	}

	return true
}
