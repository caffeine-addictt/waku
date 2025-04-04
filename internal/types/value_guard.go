package types

import (
	"errors"
	"fmt"

	"github.com/caffeine-addictt/waku/internal/config"
)

// A value guard to validate values being set
type ValueGuard[T any] struct {
	// The value
	value T
	// The validation and mutation function
	parseValue func(v T) (T, error)
	// The human readable string type
	typeString string
}

// Creating a new value guard with value parsing
func TryNewValueGuard[T any](value T, parseValue func(v T) (T, error), typeString string) (*ValueGuard[T], error) {
	v := ValueGuard[T]{
		parseValue: parseValue,
		typeString: typeString,
	}

	if err := v.Set(value); err != nil {
		return nil, err
	}
	return &v, nil
}

// Creating a new value guard without validating value
func NewValueGuard[T any](value T, parseValue func(v T) (T, error), typeString string) *ValueGuard[T] {
	return &ValueGuard[T]{
		value:      value,
		typeString: typeString,
		parseValue: parseValue,
	}
}

// Creating a new value guard without value parsing
func NewValueGuardNoParsing[T any](value T, typeString string) *ValueGuard[T] {
	return &ValueGuard[T]{
		value:      value,
		typeString: typeString,
	}
}

// Access the underlying value
func (v *ValueGuard[T]) Value() T {
	return v.value
}

// Mutate the underlying value
func (v *ValueGuard[T]) Set(value T) error {
	if v.parseValue != nil {
		value, err := v.parseValue(value)
		if err != nil {
			return errors.Join(fmt.Errorf("cannot set value %v", value), err)
		}

		v.value = value
		return nil
	}

	v.value = value
	return nil
}

// Access the human readable string type
func (v *ValueGuard[T]) Type() string {
	return v.typeString
}

// Return the string representation
func (v *ValueGuard[T]) String() string {
	return fmt.Sprintf("%v", v.value)
}

func (v *ValueGuard[T]) marshal(cfg config.ConfigType) ([]byte, error) { return cfg.Marshal(v.value) }
func (v *ValueGuard[T]) unmarshal(cfg config.ConfigType, data []byte) error {
	var tmp T
	if err := cfg.Unmarshal(data, &tmp); err != nil {
		return err
	}

	return v.Set(tmp)
}

func (v *ValueGuard[T]) UnmarshalJSON(data []byte) error {
	return v.unmarshal(config.JsonConfig{}, data)
}
func (v ValueGuard[T]) MarshalJSON() ([]byte, error) { return v.marshal(config.JsonConfig{}) }

func (v *ValueGuard[T]) UnmarshalYAML(data []byte) error {
	return v.unmarshal(config.YamlConfig{}, data)
}
func (v ValueGuard[T]) MarshalYAML() ([]byte, error) { return v.marshal(config.YamlConfig{}) }
