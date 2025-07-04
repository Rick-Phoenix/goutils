package u

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
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
