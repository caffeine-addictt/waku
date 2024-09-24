package types_test

import (
	"testing"

	"github.com/caffeine-addictt/waku/cmd/utils/types"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

func TestRegexString_UnmarshalJSON(t *testing.T) {
	tt := []struct {
		input         string
		expectedRegex string
		rule          string
		expectError   bool
	}{
		{
			input:         `"^a.*b$"`,
			expectedRegex: "^a.*b$",
			rule:          "valid regex",
			expectError:   false,
		},
		{
			input:         `"\\d{3}-\\d{2}-\\d{4}"`,
			expectedRegex: "\\d{3}-\\d{2}-\\d{4}",
			rule:          "Valid regex for SSN",
			expectError:   false,
		},
		{
			input:         `"invalid(regex"`,
			expectedRegex: "",
			rule:          "invalid regex",
			expectError:   true,
		},
		{
			input:         `"^$"`,
			expectedRegex: "^$",
			rule:          "valid regex for empty string",
			expectError:   false,
		},
		{
			input:         `"(?i)abc"`,
			expectedRegex: "(?i)abc",
			rule:          "valid regex with case insensitive flag",
			expectError:   false,
		},
		{
			input:         `"^@?(.*?)\\s*$"`,
			expectedRegex: "^@?(.*?)\\s*$",
			rule:          "valid regex with optional leading @",
			expectError:   false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.rule, func(t *testing.T) {
			var r types.RegexString
			err := json.Unmarshal([]byte(tc.input), &r)

			if tc.expectError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedRegex, r.String())
		})
	}
}
