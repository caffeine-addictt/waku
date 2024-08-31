package types_test

import (
	"testing"

	"github.com/caffeine-addictt/template/cmd/utils/types"
	"github.com/stretchr/testify/assert"
)

func TestHexColors(t *testing.T) {
	tt := []struct {
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

	for _, tc := range tt {
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
