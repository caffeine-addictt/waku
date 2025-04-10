package config

import (
	"path"
	"path/filepath"

	"github.com/caffeine-addictt/waku/internal/errors"
	"github.com/caffeine-addictt/waku/internal/types"
	"github.com/caffeine-addictt/waku/internal/utils"
)

type (
	TemplateStyles map[types.CleanString]TemplateStyle

	TemplateStyle struct {
		Ignore    *TemplateIgnore   `json:"ignore,omitempty" yaml:"ignore,omitempty"`       // The files that should be ignored when copying
		Source    types.CleanString `json:"source" yaml:"source"`                           // The source template path
		Labels    TemplateLabel     `json:"labels,omitempty" yaml:"labels,omitempty"`       // The repository labels
		Prompts   TemplatePrompts   `json:"prompts,omitempty" yaml:"prompts,omitempty"`     // The additional prompts to use
		Includes  TemplateIncludes  `json:"includes,omitempty" yaml:"includes,omitempty"`   // The additional includes
		Variables TemplateVariables `json:"variables,omitempty" yaml:"variables,omitempty"` // The additional variables
	}
)

func (t *TemplateStyles) Validate(templateRootDir string) error {
	for name, style := range *t {
		if !filepath.IsLocal(style.Source.String()) {
			return errors.
				NewWakuErrorf("path is not local: %s", style.Source).
				WithMeta("style", name.String())
		}

		styleSourceDir := path.Join(templateRootDir, style.Source.String())
		if styleSourceDir == "." {
			return errors.NewWakuErrorf("cannot use . as a path").WithMeta("style", name.String())
		}

		ok, err := utils.IsDir(styleSourceDir)
		if err != nil {
			return err
		}

		if !ok {
			return errors.NewWakuErrorf("not a directory: %s", styleSourceDir).WithMeta("style", name.String())
		}

		// Others
		if style.Ignore != nil {
			if err := style.Ignore.Validate(styleSourceDir); err != nil {
				return errors.ToWakuError(err).WithMeta("style", name.String()).WithMeta("field", "ignore")
			}
		}
		if style.Includes != nil {
			if err := style.Includes.Validate(templateRootDir, styleSourceDir); err != nil {
				return errors.ToWakuError(err).WithMeta("style", name.String()).WithMeta("field", "include")
			}
		}
	}

	return nil
}
