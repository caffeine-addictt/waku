package config

import (
	"fmt"
	"strings"

	"github.com/caffeine-addictt/template/cmd/options"
	"github.com/caffeine-addictt/template/cmd/utils/types"
)

// TemplatePrompts are the additional things that are formatted
// into the template.
//
// We take the key, strip leading/trailing whitespace and turn it to UPPER.
// The value is used to ask the user for the value.
//
// I.e.
//
//	json`{"prompts": {"my_key": "my_value"}}`
//	`aaaaa{{MY_KEY}}bbbbb` -> `aaaaamy_valuebbbb`
type TemplatePrompts map[types.CleanString]types.CleanString

func (t TemplatePrompts) Validate() error {
	keys := make([]types.CleanString, 0, len(t))
	for k := range t {
		keys = append(keys, k)
	}

	for _, k := range keys {
		newK := strings.TrimSpace(k.String())
		if newK == "" {
			return fmt.Errorf("extra template variable is empty")
		}

		v := t[k]
		newV := strings.TrimSpace(v.String())
		if newV == "" {
			newV = fmt.Sprintf("Value for '%s' template variable?", newK)
			options.Debugf("'%s' template variable ASK is empty, using defaults\n", newK)
		}

		delete(t, k)
		t[types.CleanString(strings.ToUpper(newK))] = types.CleanString(newV)
	}

	return nil
}
