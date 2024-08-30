package config

// The template.json file
type TemplateJson struct {
	Setup  *TemplateSetup  `json:"setup,omitempty"`  // Paths to executable files for post-setup
	Ignore *TemplateIgnore `json:"ignore,omitempty"` // The files that should be ignored when copying
	Labels *TemplateLabel  `json:"labels,omitempty"` // The repository labels
	Styles *TemplateStyles `json:"styles,omitempty"` // The name of the style mapped to the path to the direcotry
	Skip   *TemplateSteps  `json:"skip,omitempty"`   // The setps to skip in using the template
	Name   string          `json:"name,omitempty"`   // The name of the template
}

func (t *TemplateJson) Validate(root string) error {
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

	return nil
}
