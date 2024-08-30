package types_test

import (
	"testing"

	"github.com/caffeine-addictt/template/cmd/utils/types"
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
			if err != nil && !tc.errors {
				t.Errorf("%v. expected to error but got %s: %v", tc.rule, c, err)
			}
			if err == nil && tc.errors {
				t.Errorf("%v. expected to not error but got %s: %v", tc.rule, c, err)
			}
		})
	}
}
