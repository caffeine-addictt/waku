package types

import (
	"fmt"
	"regexp"

	"github.com/goccy/go-json"
	"gopkg.in/yaml.v3"
)

type HexColor string

var hexColorRegex = regexp.MustCompile(`^#(?:[0-9a-fA-F]{3}){1,2}$`)

func (c *HexColor) UnmarshalJSON(data []byte) error {
	var color string

	if err := json.Unmarshal(data, &color); err != nil {
		return err
	}

	if !hexColorRegex.MatchString(string(color)) {
		return fmt.Errorf("invalid hex color: %s", color)
	}

	*c = HexColor(color)
	return nil
}

func (c *HexColor) UnmarshalYAML(node *yaml.Node) error {
	var color string
	if err := node.Decode(&color); err != nil {
		return err
	}

	if !hexColorRegex.MatchString(string(color)) {
		return fmt.Errorf("invalid hex color: %s", color)
	}

	*c = HexColor(color)
	return nil
}
