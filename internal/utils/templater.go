package utils

import (
	"bufio"
	"context"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/goccy/go-json"
)

func take2NumArgs(f func(float64, float64) float64) func(string, string) string {
	return func(a, b string) string {
		aF, _ := strconv.ParseFloat(a, 64)
		bF, _ := strconv.ParseFloat(b, 64)

		return fmt.Sprintf("%v", f(aF, bF))
	}
}

var funcMap = template.FuncMap{
	// String
	"toLower":   strings.ToLower,
	"tpUpper":   strings.ToUpper,
	"toTitle":   strings.ToTitle,
	"trim":      strings.TrimSpace,
	"replace":   strings.ReplaceAll,
	"contains":  strings.Contains,
	"hasPrefix": strings.HasPrefix,
	"hasSuffix": strings.HasSuffix,
	"join":      strings.Join,
	"split":     strings.Split,
	"slug": func(s string) string {
		s = strings.ToLower(s)
		return strings.ReplaceAll(s, " ", "-")
	},

	// Arrhythmic
	"add": take2NumArgs(func(a, b float64) float64 { return a + b }),
	"sub": take2NumArgs(func(a, b float64) float64 { return a - b }),
	"mul": take2NumArgs(func(a, b float64) float64 { return a * b }),
	"div": take2NumArgs(func(a, b float64) float64 { return a / b }), // no panic IEEE-754
	"ternary": func(cond bool, a, b any) any {
		if cond {
			return a
		}
		return b
	},

	"timefmt": func(t time.Time, layout string) string {
		return t.Format(layout)
	},

	"json": func(v any) string {
		b, _ := json.MarshalIndent(v, "", "  ")
		return string(b)
	},

	"default": func(a, b string) string {
		if a == "" {
			return b
		}
		return a
	},
}

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

	t, err := template.New("file").Option("missingkey=error").Delims("{{{", "}}}").Funcs(funcMap).Parse(s.String())
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
