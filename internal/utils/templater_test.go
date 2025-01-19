package utils_test

import (
	"bufio"
	"bytes"
	"context"
	"testing"

	"github.com/caffeine-addictt/waku/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseTemplateFile(t *testing.T) {
	tests := []struct {
		name   string
		tmpl   map[string]any
		input  string
		output string
	}{
		{
			name: "Basic replacement",
			tmpl: map[string]any{
				"Name": "John",
			},
			input:  "Hello {{{ .Name }}}, welcome!",
			output: "Hello John, welcome!\n",
		},
		{
			name: "Multiple replacements",
			tmpl: map[string]any{
				"Name":  "John",
				"Place": "office",
			},
			input:  "Hello {{{ .Name }}}, welcome to the {{{ .Place }}}.",
			output: "Hello John, welcome to the office.\n",
		},
		{
			name: "No replacement",
			tmpl: map[string]any{
				"Name": "John",
			},
			input:  "No template here.",
			output: "No template here.\n",
		},
		{
			name: "Empty input",
			tmpl: map[string]any{
				"Name": "John",
			},
			input:  "",
			output: "",
		},
		{
			name: "Special characters in template key",
			tmpl: map[string]any{
				"Url": "https://example.com",
			},
			input:  "Visit {{{ \"{{\" }}}{{{ .Url }}}{{{ \"}}\" }}} for more info.",
			output: "Visit {{https://example.com}} for more info.\n",
		},
		{
			name: "Invalid template with no match",
			tmpl: map[string]any{
				"Name": "John",
			},
			input:  "Hello {{{ .Username }}}, welcome!",
			output: "Hello , welcome!\n",
		},
		{
			name: "Valid template with multiple lines",
			tmpl: map[string]any{
				"Name": "John",
			},
			input:  "Hello {{{ .Name }}}, welcome!\nHello {{{ .Name }}}, welcome!",
			output: "Hello John, welcome!\nHello John, welcome!\n",
		},
		{
			name:   "Ensure comments are not removed",
			tmpl:   map[string]any{},
			input:  "Hihi <!-- this is a comment -->\n",
			output: "Hihi <!-- this is a comment -->\n",
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewScanner(bytes.NewReader([]byte(tt.input)))
			var output bytes.Buffer
			writer := bufio.NewWriter(&output)

			err := utils.ParseTemplateFile(ctx, tt.tmpl, reader, writer)
			assert.NoError(t, err, "failed to parse")

			writer.Flush()
			assert.Equal(t, tt.output, output.String(), "wrong output")
		})
	}
}

func BenchmarkParseTemplateFile(b *testing.B) {
	tmpl := map[string]any{
		"Name":    "John",
		"Place":   "office",
		"Url":     "https://example.com",
		"Project": "my proj",
	}

	// Input strings to test different cases
	input := "Hello {{ .Name }}, welcome to the {{{ .Place }}}. For more details, visit {{{ .Url }}}. Your project is {{{ .Project }}}."
	inputComplex := "Hello {{{{ .Name }}}}. Are you visiting the {{{ .Place }}}? The details are on {{{ .Url }}}. {{{ .Project }}} is running well."

	// Benchmark different input sizes and complexities
	tt := []struct {
		name  string
		input string
	}{
		{"Small input", "Hello {{{ .Name }}}"},
		{"Medium input", input},
		{"Large input", input + input + input},
		{"Complex input", inputComplex},
	}

	ctx := context.Background()
	for _, tc := range tt {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				reader := bufio.NewScanner(bytes.NewReader([]byte(tc.input)))
				var output bytes.Buffer
				writer := bufio.NewWriter(&output)

				err := utils.ParseTemplateFile(ctx, tmpl, reader, writer)
				if err != nil {
					b.Fatalf("ParseTemplateFile() error = %v", err)
				}

				writer.Flush()
			}
		})
	}
}
