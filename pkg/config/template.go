package config

import (
	"fmt"
)

// The config file
type TemplateJson struct {
	Ignore  *TemplateIgnore `json:"ignore,omitempty" yaml:"ignore,omitempty"`   // The files that should be ignored when copying
	Labels  TemplateLabel   `json:"labels,omitempty" yaml:"labels,omitempty"`   // The repository labels
	Styles  TemplateStyles  `json:"styles" yaml:"styles"`                       // The name of the style mapped to the path to the directory
	Prompts TemplatePrompts `json:"prompts,omitempty" yaml:"prompts,omitempty"` // The additional prompts to use
}

func (t *TemplateJson) Validate(root string) error {
	if len(t.Styles) == 0 {
		return fmt.Errorf("'styles' cannot be empty")
	}

	if t.Ignore != nil {
		if err := t.Ignore.Validate(root); err != nil {
			return err
		}
	}
	if t.Styles != nil {
		if err := t.Styles.Validate(root); err != nil {
			return err
		}
	}

	return nil
}
