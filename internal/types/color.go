package types

import (
	"fmt"
	"regexp"

	"github.com/caffeine-addictt/waku/internal/config"
)

type HexColor string

var hexColorRegex = regexp.MustCompile(`^#(?:[0-9a-fA-F]{3}){1,2}$`)

func (c *HexColor) unmarshal(cfg config.ConfigType, data []byte) error {
	var color string

	if err := cfg.Unmarshal(data, &color); err != nil {
		return err
	}

	if !hexColorRegex.MatchString(string(color)) {
		return fmt.Errorf("invalid hex color: %s", color)
	}

	*c = HexColor(color)
	return nil
}

func (c *HexColor) UnmarshalYAML(data []byte) error {
	return c.unmarshal(config.YamlConfig{}, data)
}

func (c *HexColor) UnmarshalJSON(data []byte) error {
	return c.unmarshal(config.JsonConfig{}, data)
}
