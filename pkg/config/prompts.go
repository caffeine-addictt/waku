package config

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/caffeine-addictt/waku/internal/config"
	"github.com/caffeine-addictt/waku/internal/types"
	"github.com/charmbracelet/huh"
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

type (
	// TemplatePrompts are the additional things that are formatted
	// into the template.
	TemplatePrompts []TemplatePrompt

	// TemplatePrompts are the additional things that are formatted
	// into the template.
	//
	// They prompt keys are case sensitive
	// and Pacal case is recommended.
	TemplatePrompt struct {
		value     any
		Format    *string                `json:"fmt,omitempty" yaml:"fmt,omitempty"`
		Separator *string                `json:"sep,omitempty" yaml:"sep,omitempty"`
		Capture   *types.RegexString     `json:"capture,omitempty" yaml:"capture,omitempty"`
		Validate  *types.RegexString     `json:"validate,omitempty" yaml:"validate,omitempty"`
		Key       types.CleanString      `json:"key" yaml:"key"`
		Ask       types.PermissiveString `json:"ask,omitempty" yaml:"ask,omitempty"`
		Type      TemplatePromptType     `json:"type" yaml:"type"`
	}

	mockTemplatePrompt TemplatePrompt
)

// FormattedAsk returns the formatted string for the prompt
func (t *TemplatePrompt) FormattedAsk() string {
	s := string(t.Ask)

	if s == "" {
		s = t.Key.String()
		if !strings.HasSuffix(s, "?") {
			s += "?"
		}
	}

	if t.Type == TemplatePromptTypeArray {
		s += fmt.Sprintf(" [separated by '%s']", *t.Separator)
	}

	return s
}

func (t *TemplatePrompt) GetPrompt(f map[string]any) *huh.Text {
	return huh.NewText().Title(t.FormattedAsk()).Validate(func(s string) error {
		if err := t.Set(s); err != nil {
			return err
		}

		f[t.Key.String()] = t.value
		return nil
	})
}

// Set sets the value provided by the user
func (t *TemplatePrompt) Set(s string) error {
	switch t.Type {
	case TemplatePromptTypeString:
		val, err := t.formatValue(s)
		if err != nil {
			return err
		}

		t.value = val

	case TemplatePromptTypeArray:
		vals := strings.Split(s, *t.Separator)
		for i, v := range vals {
			val, err := t.formatValue(v)
			if err != nil {
				return err
			}

			vals[i] = val
		}

		t.value = vals

	default:
		panic(fmt.Sprintf("unexpected prompt type while setting value: %s", t.Type))
	}

	return nil
}

func (t *TemplatePrompt) formatValue(val string) (string, error) {
	matches := t.Capture.FindStringSubmatch(val)
	if len(matches) < 2 {
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

func (t *TemplatePrompt) unmarshal(cfg config.ConfigType, data []byte) error {
	var mock mockTemplatePrompt
	if err := cfg.Unmarshal(data, &mock); err != nil {
		return err
	}

	tp := TemplatePrompt(mock)

	// type
	tp.Type = TemplatePromptType(strings.ToLower(string(tp.Type)))
	switch tp.Type {
	case TemplatePromptTypeString, TemplatePromptTypeArray:
	default:
		return fmt.Errorf("%s is not a valid prompt type", tp.Type)
	}

	// sep
	if tp.Separator == nil {
		d := string(DefaultTemplatePromptSeparator)
		tp.Separator = &d
	}

	// capture
	if tp.Capture == nil {
		tp.Capture = &DefaultTemplatePromptCapture
	} else if tp.Capture.NumSubexp() != 1 {
		return fmt.Errorf("capture %s must have 1 sub-expression", tp.Capture.String())
	}

	// format
	if tp.Format == nil {
		d := string(DefaultTemplatePromptFormat)
		tp.Format = &d
	} else if strings.Count(*tp.Format, "*")-strings.Count(*tp.Format, "\\*") < 1 {
		return fmt.Errorf("fmt value '%s' must have at least 1 *", *tp.Format)
	}

	// validate
	if tp.Validate == nil {
		tp.Validate = &DefaultTemplatePromptValidate
	}

	*t = tp
	return nil
}

func (t *TemplatePrompt) UnmarshalYAML(data []byte) error {
	return t.unmarshal(config.YamlConfig{}, data)
}

func (t *TemplatePrompt) UnmarshalJSON(data []byte) error {
	return t.unmarshal(config.JsonConfig{}, data)
}
