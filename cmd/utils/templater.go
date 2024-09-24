package utils

import (
	"bufio"
	"context"
	"html/template"
	"strings"
)

// This handles consuming a file stream,
// templating it and writing it to writer.
//
// tmpl is a mapping of `find`: `replace`.
//
// `find` should be X where the actual template searched for is "{{X}}" {{{x}}}
//
// Time complexity through the roof: HAVE TO OPTIMIZE
func ParseTemplateFile(ctx context.Context, tmpl map[string]any, reader *bufio.Scanner, writer *bufio.Writer) error {
	var s strings.Builder
	for reader.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		s.Write(reader.Bytes())
		s.WriteRune('\n')
	}

	t, err := template.New("file").Delims("{{{", "}}}").Parse(s.String())
	if err != nil {
		return err
	}

	if err := t.Execute(writer, tmpl); err != nil {
		return err
	}

	return writer.Flush()
}

// ParseLicenseText handles templating license text
func ParseLicenseText(tmpl map[string]string, s string) string {
	for k, v := range tmpl {
		s = strings.ReplaceAll(s, `[`+k+`]`, v)
	}

	return s
}
