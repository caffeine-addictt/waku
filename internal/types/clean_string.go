package types

import (
	"fmt"
	"strings"

	"github.com/caffeine-addictt/waku/internal/config"
	"github.com/caffeine-addictt/waku/internal/utils"
)

// String that has a Clean method that also is invoked on UnmarshalJSON
type CleanString string

// Trims the string and cleans it
func (s *CleanString) Clean() {
	*s = CleanString(utils.CleanStringStrict(strings.TrimSpace(string(*s))))
}

func (s *CleanString) String() string {
	return string(*s)
}

func (s *CleanString) unmarshal(cfg config.ConfigType, data []byte) error {
	var tmp string
	if err := cfg.Unmarshal(data, &tmp); err != nil {
		return err
	}

	tmp = utils.CleanStringStrict(strings.TrimSpace(tmp))
	if tmp == "" {
		return fmt.Errorf("invalid string: %s", string(tmp))
	}

	*s = CleanString(tmp)
	return nil
}

func (s *CleanString) UnmarshalYAML(data []byte) error {
	return s.unmarshal(config.YamlConfig{}, data)
}

func (s *CleanString) UnmarshalJSON(data []byte) error {
	return s.unmarshal(config.JsonConfig{}, data)
}
