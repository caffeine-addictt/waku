package types_test

import (
	"testing"

	"github.com/caffeine-addictt/waku/internal/types"
	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"
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

		if tc.errors {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestCleanStringUnmarshalYAML(t *testing.T) {
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
		err := yaml.Unmarshal([]byte("\""+tc.in+"\""), &s)

		if tc.errors {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
