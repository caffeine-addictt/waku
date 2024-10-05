package types_test

import (
	"testing"

	"github.com/caffeine-addictt/waku/internal/types"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

var hexColorTT = []struct {
	in     string
	rule   string
	errors bool
}{
	{"#d", "hex too short", true},
	{"#ddd", "hex valid", false},
	{"#ffffff", "hex valid", false},
	{"#fff", "hex valid", false},
	{"#fffaaas", "hex too long", true},
	{"#ff", "hex too short", true},
	{"#ae24d2", "hex valid", false},
	{"sdwa2fw", "invalid letters", true},
}

func TestHexColorsJSON(t *testing.T) {
	for _, tc := range hexColorTT {
		t.Run(tc.in, func(t *testing.T) {
			c := types.HexColor(tc.in)

			err := c.UnmarshalJSON([]byte("\"" + tc.in + "\""))
			if tc.errors {
				assert.Error(t, err, tc.rule)
			} else {
				assert.NoError(t, err, tc.rule)
			}
		})
	}
}

func TestHexColorsYAML(t *testing.T) {
	for _, tc := range hexColorTT {
		t.Run(tc.in, func(t *testing.T) {
			c := types.HexColor(tc.in)

			err := yaml.Unmarshal([]byte("\""+tc.in+"\""), &c)
			if tc.errors {
				assert.Error(t, err, tc.rule)
			} else {
				assert.NoError(t, err, tc.rule)
			}
		})
	}
}
