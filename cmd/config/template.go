package config

import (
	"fmt"

	"github.com/caffeine-addictt/template/cmd/utils/types"
)

// The template.json file
type TemplateJson struct {
	Setup   *TemplateSetup    `json:"setup,omitempty"`   // Paths to executable files for post-setup
	Ignore  *TemplateIgnore   `json:"ignore,omitempty"`  // The files that should be ignored when copying
	Labels  *TemplateLabel    `json:"labels,omitempty"`  // The repository labels
	Styles  *TemplateStyles   `json:"styles,omitempty"`  // The name of the style mapped to the path to the direcotry
	Prompts *TemplatePrompts  `json:"prompts,omitempty"` // The additional prompts to use
	Name    types.CleanString `json:"name,omitempty"`    // The name of the template
}

func (t *TemplateJson) Validate(root string) error {
	// Ensure that `Name` is required if `Styles` is not present or empty
	// If `Styles` is present, `Name` must not be present
	if t.Styles == nil || len(*t.Styles) == 0 {
		if t.Name == "" {
			return fmt.Errorf("'name' is required when 'styles' is not present or empty")
		}
	} else {
		if t.Name != "" {
			return fmt.Errorf("'name' must not be present when 'styles' is provided")
		}
	}

	if t.Setup != nil {
		if err := t.Setup.Validate(root); err != nil {
			return err
		}
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
	if t.Prompts != nil {
		if err := t.Prompts.Validate(); err != nil {
			return err
		}
	}

	return nil
}
