package types

import (
	"regexp"

	"github.com/goccy/go-json"
)

type RegexString struct {
	*regexp.Regexp
}

func (r *RegexString) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	re, err := regexp.Compile(s)
	if err != nil {
		return err
	}

	r.Regexp = re
	return nil
}