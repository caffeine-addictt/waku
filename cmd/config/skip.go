package config

import (
	"fmt"
	"regexp"

	"github.com/goccy/go-json"
)

type TemplateStep string

const (
	License TemplateStep = "license"
)

type TemplateSteps []TemplateStep

var templateStepsRegexp = regexp.MustCompile(`^license$`)

func (t *TemplateSteps) UnmarshalJSON(data []byte) error {
	var tmp TemplateSteps
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	for _, v := range tmp {
		if !templateStepsRegexp.MatchString(string(v)) {
			return fmt.Errorf("invalid hex color: %s", tmp)
		}
	}

	*t = TemplateSteps(tmp)
	return nil
}
