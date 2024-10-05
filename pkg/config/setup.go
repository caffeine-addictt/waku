package config

import (
	"fmt"
	"path/filepath"
	"reflect"

	"github.com/caffeine-addictt/waku/internal/utils"
)

type TemplateSetupKey string

const (
	Linux   TemplateSetupKey = "linux"
	Windows TemplateSetupKey = "windows"
	Darwin  TemplateSetupKey = "darwin"
	Any     TemplateSetupKey = "*"
)

// Paths to executable files for post-setup
type TemplateSetup struct {
	Linux   string `json:"linux,omitempty" yaml:"linux,omitempty"`
	Darwin  string `json:"darwin,omitempty" yaml:"darwin,omitempty"`
	Windows string `json:"windows,omitempty" yaml:"windows,omitempty"`
	Any     string `json:"*,omitempty" yaml:"*,omitempty"`
}

func (t *TemplateSetup) Validate(root string) error {
	v := reflect.ValueOf(*t)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		pth := v.Field(i).Interface().(string)
		if pth == "" {
			continue
		}

		if !filepath.IsLocal(pth) {
			return fmt.Errorf("path is not local: %s", pth)
		}

		ok, err := utils.IsExecutableFile(filepath.Join(root, pth))
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("invalid executable file: %s", typeOfS.Field(i).Name)
		}
	}

	return nil
}
