package utils_test

import (
	"testing"

	"github.com/caffeine-addictt/waku/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestMultilineString(t *testing.T) {
	tt := []struct {
		out   string
		rule  string
		input []string
	}{
		{"Hello World\nHello World", "two lines", []string{"Hello World", "Hello World"}},
		{"Hello World", "one line", []string{"Hello World"}},
		{"", "empty", []string{}},
		{"\n\n", "empty newlines", []string{"", "", ""}},
		{"Hello World\nHello World", "newline character", []string{"Hello World\nHello World"}},
	}

	for _, tc := range tt {
		t.Run(tc.rule, func(t *testing.T) {
			got := utils.MultilineString(tc.input...)
			assert.Equal(t, tc.out, got, tc.rule)
		})
	}
}

func TestStringStartsWith(t *testing.T) {
	tt := []struct {
		inputs [2]string
		rule   string
		out    bool
	}{
		{[2]string{"start", "start"}, "values are euqal", true},
		{[2]string{"s", "stop"}, "stop starts with s", true},
		{[2]string{"sto", "stoop"}, "stoop starts with sto", true},
		{[2]string{"iiiiii", "iii"}, "iii is not iiiiii", false},
		{[2]string{"", ""}, "values are euqal", true},
		{[2]string{"", "nonempty"}, "empty prefix", true},
		{[2]string{"nonempty", ""}, "empty string", false},
		{[2]string{"hello", "hello world"}, "hello world starts with hello", true},
		{[2]string{"world", "hello world"}, "hello world does not start with world", false},
		{[2]string{"h", "H"}, "case sensitivity", false},
		{[2]string{"hello", "h"}, "h does not start with hello", false},
		{[2]string{"abcd", "abcde"}, "abcde starts with abcd", true},
		{[2]string{"test", "testing"}, "testing starts with test", true},
		{[2]string{"prefix", "suffix"}, "suffix does not start with prefix", false},
		{[2]string{"go", "gopher"}, "gopher starts with go", true},
		{[2]string{"😊", "😊😊😊"}, "unicode check", true},
	}

	for _, tc := range tt {
		t.Run(tc.rule, func(t *testing.T) {
			got := utils.StringStartsWith(tc.inputs[0], tc.inputs[1])
			assert.Equal(t, got, tc.out, tc.rule)
		})
	}
}

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
			got := utils.EscapeTermString(tc.in)
			assert.Equal(t, got, tc.out, tc.rule)
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
		{"\r\n", "\n", "clean up only \r", []rune{}},
		{"\x00", "", "clean up null bytes", []rune{}},
		{"\x01\x02\x03\x04\x05\x06\x07\x08\x0B\x0C\x0E\x0F", "", "clean up control characters", []rune{}},
		{"\x7f", "", "clean up DEL", []rune{}},
		{"\x1b[0m", "", "clean up full ANSI", []rune{}},
		{"\x1b", "", "clean up not started ANSI", []rune{}},
		{"\x1b[", "", "clean up uncontinued ANSI", []rune{}},
		{"\x1b[0mfoo", "foo", "clean up ANSI", []rune{}},
		{"\x1b[31mfoo\x1b[0m", "foo", "clean up ANSI", []rune{}},
		{"\x1b[31mfoo\nbar\x1b[0m", "foo\nbar", "clean up ANSI", []rune{}},
		{"\x1b[0;25;42mfoo\x1b[0m", "foo", "clean up ANSI", []rune{}},
		{"\x1b[0;25;42mfoo\x1b[0m", "foo", "clean up ANSI with multiple codes", []rune{}},
		{"abcde", "", "ignore extra runes passed too", []rune{'a', 'b', 'c', 'd', 'e'}},
		{"\x1b[31mfoo\x1b[0mbar", "foobar", "clean up ANSI with extra characters", []rune{}},
		{"\x1b[42mtext\x1b[0m\x1b[43mmore\x1b[0m", "textmore", "clean up ANSI with multiple sequences", []rune{}},
		{"\x1b[1mbold\x1b[22m\x1b[4munderline\x1b[24m", "boldunderline", "clean up bold and underline ANSI", []rune{}},
	}

	for _, tc := range tt {
		t.Run(tc.rule, func(t *testing.T) {
			got := utils.CleanString(tc.in, tc.extra...)
			assert.Equal(t, got, tc.out, tc.rule)
		})
	}
}

func TestCleanStringNoRegex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Alphanumeric string",
			input:    "Hello123",
			expected: "Hello123",
		},
		{
			name:     "String with spaces",
			input:    "Hello World 123",
			expected: "Hello World 123",
		},
		{
			name:     "String with special characters",
			input:    "Hello@World!123#",
			expected: "HelloWorld123",
		},
		{
			name:     "String with unicode characters",
			input:    "こんにちは世界123",
			expected: "こんにちは世界123",
		},
		{
			name:     "String with mixed valid and invalid characters",
			input:    "Hello, World! #123",
			expected: "Hello World 123",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "String with only special characters",
			input:    "@#$%^&*()",
			expected: "",
		},
		{
			name:     "String with leading and trailing spaces",
			input:    "  Hello World  ",
			expected: "  Hello World  ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.CleanStringNoRegex(tt.input)
			if result != tt.expected {
				t.Errorf("CleanStringNoRegex(%q) = %q; expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkStringStartsWith(b *testing.B) {
	cases := []struct {
		prefix string
		full   string
	}{
		{"", ""},
		{"a", "a"},
		{"hello", "hello world"},
		{"go", "gopher"},
		{"abc", "abcdefghijklmnopqrstuvwxyz"},
		{"longprefix", "longprefixandaverylongstring"},
		{"😊", "😊😊😊😊😊😊😊😊"},
	}

	for _, bc := range cases {
		b.Run(bc.prefix+"_"+bc.full, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				utils.StringStartsWith(bc.prefix, bc.full)
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
