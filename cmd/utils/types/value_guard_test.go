package types_test

import (
	"fmt"
	"testing"

	"github.com/caffeine-addictt/template/cmd/utils/types"
	"github.com/stretchr/testify/assert"
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
