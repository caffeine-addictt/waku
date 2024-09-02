package config

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/caffeine-addictt/template/cmd/utils"
	"github.com/caffeine-addictt/template/cmd/utils/types"
)

type TemplateStyles map[types.CleanString]struct {
	Setup  *TemplateSetup    `json:"setup,omitempty"`  // Paths to executable files for post-setup
	Ignore *TemplateIgnore   `json:"ignore,omitempty"` // The files that should be ignored when copying
	Labels *TemplateLabel    `json:"labels,omitempty"` // The repository labels
	Skip   *TemplateSteps    `json:"skip,omitempty"`   // The setps to skip in using the template
	Source types.CleanString `json:"source"`           // The source template path
	Name   types.CleanString `json:"name"`             // The name of the template
}

func (t *TemplateStyles) Validate(root string) error {
	for _, style := range *t {
		// Source
		if !filepath.IsLocal(style.Source.String()) {
			return fmt.Errorf("path is not local: %s", style.Source)
		}

		resolvedPath := path.Join(root, style.Source.String())
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

		// Others
		if style.Setup != nil {
			if err := style.Setup.Validate(root); err != nil {
				return err
			}
		}
		if style.Ignore != nil {
			if err := style.Ignore.Validate(root); err != nil {
				return err
			}
		}
	}

	return nil
}
