package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/caffeine-addictt/waku/internal/types"
	"gopkg.in/yaml.v3"
)

type TemplateIgnore types.Set[string]

func (t *TemplateIgnore) Validate(styleSourceDir string) error {
	for path := range *t {
		ignorePath := strings.TrimSpace(path)

		// handle bang
		ignorePath = strings.TrimPrefix(ignorePath, "!")

		// handle glob
		isGlob := false
		if strings.HasSuffix(ignorePath, "/*") {
			isGlob = true
			ignorePath = strings.TrimSuffix(ignorePath, "/*")
		}

		if !filepath.IsLocal(ignorePath) {
			return fmt.Errorf("path is not local: %s", path)
		}

		fileinfo, err := os.Stat(filepath.Join(styleSourceDir, ignorePath))
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}

		if isGlob && !fileinfo.IsDir() {
			return fmt.Errorf("%s is not a directory", path)
		}
	}

	return nil
}

// UnmarshalJSON unmarshals a JSON array into a set
func (t *TemplateIgnore) UnmarshalJSON(data []byte) error {
	var items []string
	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}
	*t = TemplateIgnore(types.NewSet(items...))
	return nil
}

// MarshalJSON marshals a set into a JSON array
func (t TemplateIgnore) MarshalJSON() ([]byte, error) {
	s := types.Set[string](t)
	return json.Marshal(s.ToSlice())
}

// UnmarshalYAML unmarshals a YAML string into a set
func (t *TemplateIgnore) UnmarshalYAML(node *yaml.Node) error {
	var items []string
	if err := node.Decode(&items); err != nil {
		return err
	}
	*t = TemplateIgnore(types.NewSet(items...))
	return nil
}

// MarshalYAML marshals a set into a YAML string
func (t TemplateIgnore) MarshalYAML() (interface{}, error) {
	s := types.Set[string](t)
	return s.ToSlice(), nil
}
