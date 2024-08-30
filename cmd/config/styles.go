package config

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/caffeine-addictt/template/cmd/utils"
)

type TemplateStyles map[string]string

func (t *TemplateStyles) Validate(root string) error {
	for _, pth := range *t {
		if !filepath.IsLocal(pth) {
			return fmt.Errorf("path is not local: %s", pth)
		}

		resolvedPath := path.Join(root, pth)
		if resolvedPath == "." {
			return fmt.Errorf("cannot use . as a path")
		}

		ok, err := utils.IsDir(resolvedPath)
		if err != nil {
			return err
		}

		if !ok {
			return fmt.Errorf("not a directory: %s", resolvedPath)
		}
	}

	return nil
}
