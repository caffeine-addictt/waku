package config

import (
	"fmt"
	"strings"

	"github.com/caffeine-addictt/waku/internal/config"
)

// The type of the prompt response
type TemplatePromptType string

const (
	TemplatePromptTypeString TemplatePromptType = "str"
	TemplatePromptTypeArray  TemplatePromptType = "arr"
)

func (t *TemplatePromptType) unmarshal(cfg config.ConfigType, data []byte) error {
	var str string
	err := cfg.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	tp := TemplatePromptType(strings.ToLower(string(str)))
	switch tp {
	case TemplatePromptTypeString, TemplatePromptTypeArray:
	default:
		return fmt.Errorf("%s is not a valid prompt type", tp)
	}

	*t = tp
	return nil
}

func (t *TemplatePromptType) UnmarshalJSON(data []byte) error {
	return t.unmarshal(config.JsonConfig{}, data)
}

func (t *TemplatePromptType) UnmarshalYAML(data []byte) error {
	return t.unmarshal(config.YamlConfig{}, data)
}
