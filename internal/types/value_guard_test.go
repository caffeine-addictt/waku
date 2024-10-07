package types_test

import (
	"fmt"
	"testing"

	"github.com/caffeine-addictt/waku/internal/types"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestNoParsing(t *testing.T) {
	val := "test"
	typeString := "my type"

	v := types.NewValueGuardNoParsing(val, typeString)
	checkValues(t, val, typeString, v)

	err := v.Set("new value")
	assert.NoError(t, err, v.String(), "failed to set value")
}

func TestParsing(t *testing.T) {
	val := "test fail"
	typeString := "my type"

	v := types.NewValueGuard(val, func(s string) (string, error) {
		if s == "test fail" {
			return "", fmt.Errorf("failed to parse")
		}
		return s, nil
	}, typeString)
	checkValues(t, val, typeString, v)

	// Only error when setting
	err := v.Set(val)
	assert.Error(t, err, "expected error")
}

func TestParsingJsonUnMarshal(t *testing.T) {
	var tmp types.ValueGuard[string]
	err := json.Unmarshal([]byte(`"test"`), &tmp)
	assert.NoError(t, err, "failed to unmarshal")
	checkValues(t, "test", "", &tmp)
}

func TestParsingJsonMarshal(t *testing.T) {
	tmp := types.NewValueGuardNoParsing("hi", "ok")
	data, err := json.Marshal(tmp)
	assert.NoError(t, err, "failed to marshal")
	assert.Equal(t, `"hi"`, string(data), "value should match")
}

func TestParsingYamlUnMarshal(t *testing.T) {
	var tmp types.ValueGuard[string]
	err := yaml.Unmarshal([]byte(`"test"`), &tmp)
	assert.NoError(t, err, "failed to unmarshal")
	checkValues(t, "test", "", &tmp)
}

func TestParsingYamlMarshal(t *testing.T) {
	tmp := types.NewValueGuardNoParsing("hi", "ok")
	data, err := yaml.Marshal(tmp)
	assert.NoError(t, err, "failed to marshal")
	assert.Equal(t, "hi\n", string(data), "value should match")
}

func TestParsingFailEarly(t *testing.T) {
	val := "test fail"
	typeString := "my type"

	_, err := types.TryNewValueGuard(val, func(s string) (string, error) {
		if s == "test fail" {
			return "", fmt.Errorf("failed to parse")
		}
		return s, nil
	}, typeString)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func checkValues(t *testing.T, val, typeString string, vg *types.ValueGuard[string]) {
	assert.Equal(t, val, vg.Value(), "value Value() should match")
	assert.Equal(t, val, vg.String(), "value String() should match")
	assert.Equal(t, typeString, vg.Type(), "value Type() should match")
}
