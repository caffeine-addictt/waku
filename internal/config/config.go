// For joining the Marshal and Unmarshal
// methods for all supported config types.
package config

import (
	"github.com/goccy/go-json"
	"github.com/goccy/go-yaml"
)

// ConfigType is a generic interface for
// marshalling and unmarshalling config types
type ConfigType interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}

// JsonConfig implements the ConfigType interface
type JsonConfig struct{}

func (JsonConfig) Unmarshal(data []byte, v interface{}) error { return json.Unmarshal(data, v) }
func (JsonConfig) Marshal(v interface{}) ([]byte, error)      { return json.Marshal(v) }

// YamlConfig implements the ConfigType interface
type YamlConfig struct{}

func (YamlConfig) Unmarshal(data []byte, v interface{}) error { return yaml.Unmarshal(data, v) }
func (YamlConfig) Marshal(v interface{}) ([]byte, error)      { return yaml.Marshal(v) }
