package types

import (
	"regexp"

	"github.com/caffeine-addictt/waku/internal/config"
)

type RegexString struct {
	*regexp.Regexp
}

func (r *RegexString) unmarshal(cfg config.ConfigType, data []byte) error {
	var s string
	if err := cfg.Unmarshal(data, &s); err != nil {
		return err
	}

	re, err := regexp.Compile(s)
	if err != nil {
		return err
	}

	r.Regexp = re
	return nil
}

func (r *RegexString) UnmarshalYAML(data []byte) error {
	return r.unmarshal(config.YamlConfig{}, data)
}

func (r *RegexString) UnmarshalJSON(data []byte) error {
	return r.unmarshal(config.JsonConfig{}, data)
}
