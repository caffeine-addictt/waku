package utils_test

import (
	"testing"

	"github.com/caffeine-addictt/template/cmd/utils"
)

func TestEscapingTermString(t *testing.T) {
	tt := []struct {
		in   string
		out  string
		rule string
	}{
		{"git | sed", "git   sed", "escape pipes"},
		{"a & b", "a   b", "escape chaining commands"},
		{"a && b", "a    b", "escape AND commands"},
		{"a || b", "a    b", "escape OR operator"},
		{"a; b", "a  b", "escape semicolons"},
		{"a`b", "a b", "escape backticks"},
		{"$b", " b", "escape dollar sign"},
		{"a\nb", "a b", "escape newlines"},
		{"a\rb", "a b", "escape carriage returns"},
		{"|&\n\r;`$", "       ", "escape all known special characters"},
	}

	for _, tc := range tt {
		t.Run(tc.rule, func(t *testing.T) {
			if got := utils.EscapeTermString(tc.in); got != tc.out {
				t.Errorf("%v. EscapeTermString() = '%v', want '%v'", tc.rule, got, tc.out)
			}
		})
	}
}
