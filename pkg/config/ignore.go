package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/caffeine-addictt/waku/internal/config"
	"github.com/caffeine-addictt/waku/internal/types"
)

type TemplateIgnore types.Set[string]

func (t *TemplateIgnore) Validate(styleSourceDir string) error {
	for path := range *t {
		ignorePath := strings.TrimSpace(path)

		// handle bang
		ignorePath = strings.TrimPrefix(ignorePath, "!")

		// handle glob
		isDir := false
		if strings.HasSuffix(ignorePath, "/*") || strings.HasSuffix(ignorePath, "/**") || strings.HasSuffix(ignorePath, "/") {
			isDir = true
			ignorePath = strings.TrimSuffix(strings.TrimSuffix(strings.TrimSuffix(ignorePath, "*"), "*"), "/")
		}

		if !filepath.IsLocal(ignorePath) {
			return fmt.Errorf("path is not local: %s", path)
		}

		// skip globs for now (something to consider implementing in the future)
		if strings.Contains(ignorePath, "*") {
			return nil
		}

		fileinfo, err := os.Stat(filepath.Join(styleSourceDir, ignorePath))
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}

		if isDir && !fileinfo.IsDir() {
			return fmt.Errorf("%s is not a directory", path)
		}
	}

	return nil
}

func (t *TemplateIgnore) unmarshall(cfg config.ConfigType, data []byte) error {
	var items []string
	if err := cfg.Unmarshal(data, &items); err != nil {
		return err
	}
	*t = TemplateIgnore(types.NewSet(items...))
	return nil
}

func (t *TemplateIgnore) marshal(cfg config.ConfigType) ([]byte, error) {
	s := types.Set[string](*t)
	return cfg.Marshal(s.ToSlice())
}

func (t *TemplateIgnore) UnmarshalJSON(data []byte) error {
	return t.unmarshall(config.JsonConfig{}, data)
}
func (t *TemplateIgnore) MarshalJSON() ([]byte, error) { return t.marshal(config.JsonConfig{}) }

func (t *TemplateIgnore) UnmarshalYAML(data []byte) error {
	return t.unmarshall(config.YamlConfig{}, data)
}
func (t *TemplateIgnore) MarshalYAML() ([]byte, error) { return t.marshal(config.YamlConfig{}) }
