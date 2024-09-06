package utils

import (
	"bufio"
	"context"
	"regexp"
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
func ParseTemplateFile(ctx context.Context, tmpl map[string]string, reader *bufio.Scanner, writer *bufio.Writer) error {
	// generate the regexp
	reg := make(map[*regexp.Regexp]string, len(tmpl))
	for find, replace := range tmpl {
		reg[regexp.MustCompile(`{\s*{\s*`+CleanStringNoRegex(find)+`\s*}\s*}`)] = replace
	}

	for reader.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		line := reader.Text()
		for find, replace := range reg {
			line = find.ReplaceAllString(line, replace)
		}

		if _, err := writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}

	return nil
}

// ParseLicenseText handles templating license text
func ParseLicenseText(tmpl map[string]string, s string) string {
	for k, v := range tmpl {
		s = strings.ReplaceAll(s, `[`+k+`]`, v)
	}

	return s
}
