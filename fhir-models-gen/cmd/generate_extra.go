package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//go:embed extra/types.go
var extraTypes string

func GenerateExtra(dir string) error {
	fpath := path.Join(dir, "extraTypes.gen.go")
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", fpath, err)
	}
	defer f.Close()

	abspath, err := filepath.Abs(dir)
	if err != nil {
		return fmt.Errorf("failed to determine absolute path for target dir: %w", err)
	}

	pkgname := path.Base(abspath)
	content := strings.Replace(extraTypes, "package types", fmt.Sprintf("package %s", pkgname), 1)

	if _, err := f.WriteString(content); err != nil {
		return fmt.Errorf("failed to write to file %s: %w", fpath, err)
	}

	return nil
}
