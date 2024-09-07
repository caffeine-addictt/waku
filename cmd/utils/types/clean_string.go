package types

import (
	"fmt"
	"strings"

	"github.com/caffeine-addictt/waku/cmd/utils"
	"github.com/goccy/go-json"
)

// String that has a Clean method that also is invoked on UnmarshalJSON
type CleanString string

// Trims the string and cleans it
func (s *CleanString) Clean() {
	*s = CleanString(utils.CleanString(strings.TrimSpace(string(*s))))
}

func (s *CleanString) String() string {
	return string(*s)
}

func (s *CleanString) UnmarshalJSON(data []byte) error {
	var tmp string
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	tmp = utils.CleanString(strings.TrimSpace(tmp))
	if tmp == "" {
		return fmt.Errorf("invalid string: %s", string(tmp))
	}

	*s = CleanString(tmp)
	return nil
}
