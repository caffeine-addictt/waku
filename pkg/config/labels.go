package config

import "github.com/caffeine-addictt/waku/internal/types"

type TemplateLabel []struct {
	Name  types.CleanString `json:"name" yaml:"name"`
	Color types.HexColor    `json:"color" yaml:"color"`
	Desc  string            `json:"description,omitempty" yaml:"description,omitempty"`
}
