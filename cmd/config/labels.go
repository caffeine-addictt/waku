package config

import "github.com/caffeine-addictt/waku/cmd/utils/types"

type TemplateLabel []struct {
	Color types.HexColor `json:"color"`
	Desc  string         `json:"description,omitempty"`
}
