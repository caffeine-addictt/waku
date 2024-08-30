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
		{"git | sed", "git  sed", "escape pipes"},
		{"a & b", "a  b", "escape chaining commands"},
		{"a && b", "a  b", "escape AND commands"},
		{"a || b", "a  b", "escape OR operator"},
		{"a; b", "a b", "escape semicolons"},
		{"a`b", "ab", "escape backticks"},
		{"$b", "b", "escape dollar sign"},
		{"a\nb", "ab", "escape newlines"},
		{"a\rb", "ab", "escape carriage returns"},
		{"|&\n\r;`$", "", "escape all known special characters"},
	}

	for _, tc := range tt {
		t.Run(tc.rule, func(t *testing.T) {
			if got := utils.EscapeTermString(tc.in); got != tc.out {
				t.Errorf("%v. EscapeTermString() = '%v', want '%v'", tc.rule, got, tc.out)
			}
		})
	}
}

func TestCleaningString(t *testing.T) {
	tt := []struct {
		in    string
		out   string
		rule  string
		extra []rune
	}{
		{"\r\n", "", "clean up newlines", []rune{}},
		{"\x00", "", "clean up null bytes", []rune{}},
		{"\x01\x02\x03\x04\x05\x06\x07\x08\x0B\x0C\x0E\x0F", "", "clean up control characters", []rune{}},
		{"\x7f", "", "clean up DEL", []rune{}},
		{"\x1b[0m", "", "clean up full ANSI", []rune{}},
		{"\x1b", "", "clean up not started ANSI", []rune{}},
		{"\x1b[", "", "clean up uncontinued ANSI", []rune{}},
		{"\x1b[0mfoo", "foo", "clean up ANSI", []rune{}},
		{"\x1b[31mfoo\x1b[0m", "foo", "clean up ANSI", []rune{}},
		{"\x1b[31mfoo\nbar\x1b[0m", "foobar", "clean up ANSI", []rune{}},
		{"\x1b[0;25;42mfoo\x1b[0m", "foo", "clean up ANSI", []rune{}},
		{"\x1b[0;25;42mfoo\x1b[0m", "foo", "clean up ANSI with multiple codes", []rune{}},
		{"abcde", "", "ignore extra runes passed too", []rune{'a', 'b', 'c', 'd', 'e'}},
		{"\x1b[31mfoo\x1b[0mbar", "foobar", "clean up ANSI with extra characters", []rune{}},
		{"\x1b[42mtext\x1b[0m\x1b[43mmore\x1b[0m", "textmore", "clean up ANSI with multiple sequences", []rune{}},
		{"\x1b[1mbold\x1b[22m\x1b[4munderline\x1b[24m", "boldunderline", "clean up bold and underline ANSI", []rune{}},
	}

	for _, tc := range tt {
		t.Run(tc.rule, func(t *testing.T) {
			if got := utils.CleanString(tc.in, tc.extra...); got != tc.out {
				t.Errorf("%v. CleanString() = '%v', want '%v'", tc.rule, got, tc.out)
			}
		})
	}
}

func BenchmarkCleaningString(b *testing.B) {
	bt := []struct {
		in    string
		extra []rune
	}{
		{"\r\n", []rune{}},
		{"\x00", []rune{}},
		{"\x01\x02\x03\x04\x05\x06\x07\x08\x0B\x0C\x0E\x0F", []rune{}},
		{"\x7f", []rune{}},
		{"\x1b[0m", []rune{}},
		{"\x1b", []rune{}},
		{"\x1b[", []rune{}},
		{"\x1b[0mfoo", []rune{}},
		{"\x1b[31mfoo\x1b[0m", []rune{}},
		{"\x1b[31mfoo\nbar\x1b[0m", []rune{}},
		{"\x1b[0;25;42mfoo\x1b[0m", []rune{}},
		{"abcde", []rune{'a', 'b', 'c', 'd', 'e'}},
		{"\x1b[31mfoo\x1b[0mbar", []rune{}},
		{"\x1b[42mtext\x1b[0m\x1b[43mmore\x1b[0m", []rune{}},
		{"\x1b[1mbold\x1b[22m\x1b[4munderline\x1b[24m", []rune{}},
	}

	for _, bc := range bt {
		b.Run(bc.in, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = utils.CleanString(bc.in, bc.extra...)
			}
		})
	}
}
