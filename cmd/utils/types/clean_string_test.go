package types_test

import (
	"testing"

	"github.com/caffeine-addictt/template/cmd/utils/types"
)

func TestCleanStringUnmarshalJSON(t *testing.T) {
	tt := []struct {
		in     string
		errors bool
	}{
		{"aa", false},
		{"\r\b\n\t", true},
		{"", true},
		{" ", true},
		{"\r\ns", false},
	}

	for _, tc := range tt {
		var s types.CleanString
		err := s.UnmarshalJSON([]byte("\"" + tc.in + "\""))

		if tc.errors && err == nil {
			t.Errorf("expected error, but got nil")
		} else if !tc.errors && err != nil {
			t.Errorf("did not expect error, but got %v", err)
		}
	}
}
