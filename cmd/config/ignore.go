package config

import (
	"fmt"
	"os"
	"strings"
)

type TemplateIgnore []string

func (t *TemplateIgnore) Validate(root string) error {
	for _, path := range *t {
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
