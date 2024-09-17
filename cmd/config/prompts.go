package config

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/caffeine-addictt/waku/cmd/utils/types"
	"github.com/charmbracelet/huh"
	"github.com/goccy/go-json"
)

// The type of the prompt response
type TemplatePromptType string

const (
	TemplatePromptTypeString TemplatePromptType = "str"
	TemplatePromptTypeArray  TemplatePromptType = "arr"

	DefaultTemplatePromptSeparator string = " "
	DefaultTemplatePromptFormat    string = "*"
)

var (
	DefaultTemplatePromptCapture  types.RegexString = types.RegexString{Regexp: regexp.MustCompile(`\s*(.*?)\s*`)}
	DefaultTemplatePromptValidate types.RegexString = types.RegexString{Regexp: regexp.MustCompile(`.+`)}
)

// TemplatePrompts are the additional things that are formatted
// into the template.
type TemplatePrompts []TemplatePrompt

// TemplatePrompts are the additional things that are formatted
// into the template.
//
// They prompt keys are case sensitive
// and Pacal case is recommended.
type TemplatePrompt struct {
	Value     any
	Format    *string            `json:"fmt,omitempty"`
	Separator *string            `json:"sep,omitempty"`
	Capture   *types.RegexString `json:"capture,omitempty"`
	Validate  *types.RegexString `json:"validate,omitempty"`
	Key       types.CleanString  `json:"key"`
	Ask       types.CleanString  `json:"ask,omitempty"`
	Type      TemplatePromptType `json:"type"`
}

// FormattedAsk returns the formatted string for the prompt
func (t *TemplatePrompt) FormattedAsk() string {
	s := string(t.Ask)

	if s == "" {
		s = t.Key.String()
	}

	if !strings.HasSuffix(s, "?") {
		s += "?"
	}

	if t.Type == TemplatePromptTypeArray {
		s += fmt.Sprintf(" [separated by '%s']", *t.Separator)
	}

	return s
}

func (t *TemplatePrompt) GetPrompt() *huh.Text {
	return huh.NewText().Title(t.FormattedAsk()).Validate(t.Set)
}

// Set sets the value provided by the user
func (t *TemplatePrompt) Set(s string) error {
	switch t.Type {
	case TemplatePromptTypeString:
		val, err := t.formatValue(s)
		if err != nil {
			return err
		}

		t.Value = val

	case TemplatePromptTypeArray:
		vals := strings.Split(s, *t.Separator)
		for i, v := range vals {
			val, err := t.formatValue(v)
			if err != nil {
				return err
			}

			vals[i] = val
		}

		t.Value = vals

	default:
		panic(fmt.Sprintf("unexpected prompt type while setting value: %s", t.Type))
	}

	return nil
}

func (t *TemplatePrompt) formatValue(val string) (string, error) {
	matches := t.Capture.FindStringSubmatch(val)
	if matches == nil || len(matches) < 2 {
		return "", fmt.Errorf("capture %s did not match '%s'", t.Capture.String(), val)
	}

	var s strings.Builder
	i := 0
	for i < len(*t.Format) {
		switch (*t.Format)[i] {
		case '\\':
			if i+1 < len(*t.Format) && (*t.Format)[i+1] == '*' {
				s.WriteRune('*')
				i += 2
				continue
			}
		case '*':
			s.WriteString(val)
		default:
			s.WriteByte((*t.Format)[i])
		}

		i++
	}

	l := s.String()
	if !t.Validate.MatchString(l) {
		return "", fmt.Errorf("value '%s' did not match '%s'", l, *t.Validate)
	}

	return l, nil
}

func (t *TemplatePrompt) UnmarshalJSON(data []byte) error {
	type Alias TemplatePrompt
	var s Alias

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	// type
	s.Type = TemplatePromptType(strings.ToLower(string(s.Type)))
	switch s.Type {
	case TemplatePromptTypeString, TemplatePromptTypeArray:
	default:
		return fmt.Errorf("%s is not a valid prompt type", s.Type)
	}

	// sep
	if s.Separator == nil {
		d := string(DefaultTemplatePromptSeparator)
		s.Separator = &d
	}

	// capture
	if s.Capture == nil {
		s.Capture = &DefaultTemplatePromptCapture
	} else if s.Capture.NumSubexp() != 1 {
		return fmt.Errorf("capture %s must have 1 sub-expression", s.Capture.String())
	}

	// format
	if s.Format == nil {
		d := string(DefaultTemplatePromptFormat)
		s.Format = &d
	} else if strings.Count(*s.Format, "*")-strings.Count(*s.Format, "\\*") < 1 {
		return fmt.Errorf("fmt value '%s' must have at least 1 *", *s.Format)
	}

	// validate
	if s.Validate == nil {
		s.Validate = &DefaultTemplatePromptValidate
	}

	*t = TemplatePrompt(s)
	return nil
}
