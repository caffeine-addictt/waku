package types

import (
	"fmt"
	"strings"

	"github.com/caffeine-addictt/waku/internal/config"
	"github.com/caffeine-addictt/waku/internal/utils"
)

// String that has a Clean method that also is invoked on UnmarshalJSON
type PermissiveString string

// Trims the string and cleans it
func (s *PermissiveString) Clean() {
	*s = PermissiveString(utils.CleanString(strings.TrimSpace(string(*s))))
}

func (s *PermissiveString) String() string {
	return string(*s)
}

func (s *PermissiveString) unmarshal(cfg config.ConfigType, data []byte) error {
	var tmp string
	if err := cfg.Unmarshal(data, &tmp); err != nil {
		return err
	}

	tmp = utils.CleanString(strings.TrimSpace(tmp))
	if tmp == "" {
		return fmt.Errorf("invalid string: %s", string(tmp))
	}

	*s = PermissiveString(tmp)
	return nil
}

func (s *PermissiveString) UnmarshalYAML(data []byte) error {
	return s.unmarshal(config.YamlConfig{}, data)
}

func (s *PermissiveString) UnmarshalJSON(data []byte) error {
	return s.unmarshal(config.JsonConfig{}, data)
}
