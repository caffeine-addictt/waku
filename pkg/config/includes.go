package config

import (
	stderrors "errors"
	"path/filepath"
	"strings"

	"github.com/caffeine-addictt/waku/internal/config"
	"github.com/caffeine-addictt/waku/internal/errors"
	"github.com/caffeine-addictt/waku/internal/types"
	"github.com/caffeine-addictt/waku/internal/utils"
)

type (
	TemplateIncludes []TemplateInclude

	// Additional directories to include in templating.
	//
	// This can support []{ source, ignore } and []string
	TemplateInclude struct {
		// The files that should be ignored when copying
		Ignore *TemplateIgnore `json:"ignore,omitempty" yaml:"ignore,omitempty"`

		// The templated parent directory path
		Directory *types.CleanString `json:"dir,omitempty" yaml:"dir,omitempty"`

		// Must be:
		// - subdir of the template root
		// - not same as styleSourceDir
		// - not subdir of styleSourceDir
		Source types.CleanString `json:"source" yaml:"source"`
	}

	mockTemplateInclude struct {
		Ignore    *TemplateIgnore    `json:"ignore,omitempty" yaml:"ignore,omitempty"`
		Directory *types.CleanString `json:"dir,omitempty" yaml:"dir,omitempty"`
		Source    types.CleanString  `json:"source" yaml:"source"`
	}
)

func (t *TemplateIncludes) Validate(templateRootDir, styleSourceDir string) error {
	for _, include := range *t {
		srcPath := include.Source.String()
		if !filepath.IsLocal(srcPath) {
			return errors.NewWakuErrorf("path is not local: %s", srcPath)
		}

		srcPath = filepath.Join(templateRootDir, srcPath)
		if srcPath == "." {
			return errors.NewWakuErrorf("cannot use . as a path")
		}

		if srcPath == styleSourceDir {
			return errors.NewWakuErrorf("cannot use same path as style source: %s", srcPath)
		}

		if strings.HasPrefix(srcPath, styleSourceDir) {
			return errors.NewWakuErrorf("cannot use subdir of style source: %s", srcPath)
		}

		ok, err := utils.IsDir(srcPath)
		if err != nil {
			return err
		}

		if !ok {
			return errors.NewWakuErrorf("not a directory: %s", srcPath)
		}

		// dir
		if include.Directory != nil && !filepath.IsLocal(include.Directory.String()) {
			return errors.NewWakuErrorf("path is not local: %s", include.Directory)
		}

		// ignore
		if include.Ignore == nil {
			continue
		}
		if err := include.Ignore.Validate(styleSourceDir); err != nil {
			return err
		}
	}

	return nil
}

func (t *TemplateInclude) unmarshal(cfg config.ConfigType, data []byte) error {
	var ti mockTemplateInclude
	err := cfg.Unmarshal(data, &ti)
	if err != nil {
		var tiAlt types.CleanString
		err2 := cfg.Unmarshal(data, &tiAlt)
		if err2 != nil {
			return stderrors.Join(err2, err)
		}

		ti.Source = tiAlt
	}

	*t = TemplateInclude(ti)
	return nil
}

func (t *TemplateInclude) marshal(cfg config.ConfigType) ([]byte, error) {
	if t.Ignore == nil {
		return cfg.Marshal(t.Source)
	}
	return cfg.Marshal(mockTemplateInclude(*t))
}

func (t *TemplateInclude) UnmarshalJSON(data []byte) error {
	return t.unmarshal(config.JsonConfig{}, data)
}
func (t TemplateInclude) MarshalJSON() ([]byte, error) { return t.marshal(config.JsonConfig{}) }

func (t *TemplateInclude) UnmarshalYAML(data []byte) error {
	return t.unmarshal(config.YamlConfig{}, data)
}
func (t TemplateInclude) MarshalYAML() ([]byte, error) { return t.marshal(config.YamlConfig{}) }
