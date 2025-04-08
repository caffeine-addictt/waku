package config

import (
	"fmt"
	"strings"

	"github.com/caffeine-addictt/waku/internal/config"
)

// The type of the prompt response
type TemplateVarType string

const (
	TemplateVarTypeString TemplateVarType = "str"
	TemplateVarTypeArray  TemplateVarType = "arr"
)

func (t *TemplateVarType) unmarshal(cfg config.ConfigType, data []byte) error {
	var str string
	err := cfg.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	tp := TemplateVarType(strings.ToLower(string(str)))
	switch tp {
	case TemplateVarTypeString, TemplateVarTypeArray:
	default:
		return fmt.Errorf("%s is not a valid prompt type", tp)
	}

	*t = tp
	return nil
}

func (t *TemplateVarType) UnmarshalJSON(data []byte) error {
	return t.unmarshal(config.JsonConfig{}, data)
}

func (t *TemplateVarType) UnmarshalYAML(data []byte) error {
	return t.unmarshal(config.YamlConfig{}, data)
}
