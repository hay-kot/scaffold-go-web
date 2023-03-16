{{- if .Scaffold.use_database -}}
// Package migrations provides a way to embed migrations into the application.
package migrations

import (
	"embed"
	"os"
	"path/filepath"
)

//go:embed all:sql
var Files embed.FS

// Write writes the embedded migrations to a temporary directory.
// It returns an error and a cleanup function. The cleanup function
// should be called when the migrations are no longer needed.
func Write(temp string) error {
	err := os.MkdirAll(temp, 0o755)
	if err != nil {
		return err
	}

	fsDir, err := Files.ReadDir("migrations")
	if err != nil {
		return err
	}

	for _, f := range fsDir {
		if f.IsDir() {
			continue
		}

		b, err := Files.ReadFile(filepath.Join("migrations", f.Name()))
		if err != nil {
			return err
		}

		err = os.WriteFile(filepath.Join(temp, f.Name()), b, 0o644)
		if err != nil {
			return err
		}
	}

	return nil
}
{{- end -}}