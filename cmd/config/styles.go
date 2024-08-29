package config

import (
	"fmt"
	"path"

	"github.com/caffeine-addictt/template/cmd/utils"
)

type TemplateStyles map[string]string

func (t *TemplateStyles) Validate(root string) error {
	for _, pth := range *t {
		resolvedPath := path.Join(root, pth)

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
