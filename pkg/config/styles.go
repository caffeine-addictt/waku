package config

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/caffeine-addictt/waku/internal/types"
	"github.com/caffeine-addictt/waku/internal/utils"
)

type TemplateStyles map[types.CleanString]TemplateStyle

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
		if style.Ignore != nil {
			if err := style.Ignore.Validate(root); err != nil {
				return err
			}
		}
	}

	return nil
}

type TemplateStyle struct {
	Setup   *TemplateSetup    `json:"setup,omitempty" yaml:"setup,omitempty"`     // Paths to executable files for post-setup
	Ignore  *TemplateIgnore   `json:"ignore,omitempty" yaml:"ignore,omitempty"`   // The files that should be ignored when copying
	Source  types.CleanString `json:"source" yaml:"source"`                       // The source template path
	Labels  TemplateLabel     `json:"labels,omitempty" yaml:"labels,omitempty"`   // The repository labels
	Prompts TemplatePrompts   `json:"prompts,omitempty" yaml:"prompts,omitempty"` // The additional prompts to use
}
