package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/caffeine-addictt/template/cmd/utils/types"
)

type TemplateIgnore types.Set[string]

func (t *TemplateIgnore) Validate(root string) error {
	for path := range *t {
		dirPath := strings.TrimSpace(path)

		// handle bang
		dirPath = strings.TrimPrefix(dirPath, "!")

		// handle glob
		isGlob := false
		if strings.HasSuffix(dirPath, "/*") {
			isGlob = true
			dirPath = strings.TrimSuffix(dirPath, "/*")
		}

		if !filepath.IsLocal(dirPath) {
			return fmt.Errorf("path is not local: %s", path)
		}

		fileinfo, err := os.Stat(dirPath)
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}

		if isGlob && !fileinfo.IsDir() {
			return fmt.Errorf("%s: exists but is not a directory", path)
		}
	}

	return nil
}
