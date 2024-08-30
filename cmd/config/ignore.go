package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type TemplateIgnore []string

func (t *TemplateIgnore) Validate(root string) error {
	for _, path := range *t {
		if !filepath.IsLocal(path) {
			return fmt.Errorf("path is not local: %s", path)
		}

		dirPath := path
		isGrep := strings.HasSuffix(path, "/*")

		// Account for path/to/files/*
		if isGrep {
			dirPath = path[:len(path)-2]
		}

		fileinfo, err := os.Stat(dirPath)
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}

		if isGrep && !fileinfo.IsDir() {
			return fmt.Errorf("%s: exists but is not a directory", path)
		}
	}

	return nil
}
