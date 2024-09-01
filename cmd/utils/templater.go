package utils

import (
	"bufio"
	"regexp"
)

// This handles consuming a file stream,
// templating it and writing it to writer.
//
// tmpl is a mapping of `find`: `replace`.
//
// `find` should be X where the actual template searched for is "{{X}}" {{{x}}}
//
// Time complexity through the roof: HAVE TO OPTIMIZE
func ParseTemplateFile(tmpl map[string]string, reader *bufio.Scanner, writer *bufio.Writer) error {
	// generate the regexp
	reg := map[*regexp.Regexp]string{}
	for find, replace := range tmpl {
		reg[regexp.MustCompile(`{\s*{\s*`+CleanStringNoRegex(find)+`\s*}\s*}`)] = replace
	}

	for reader.Scan() {
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
