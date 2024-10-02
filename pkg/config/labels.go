package config

import "github.com/caffeine-addictt/waku/internal/types"

type TemplateLabel []struct {
	Name  types.CleanString `json:"name"`
	Color types.HexColor    `json:"color"`
	Desc  string            `json:"description,omitempty"`
}
