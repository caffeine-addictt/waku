package types_test

import (
	"fmt"
	"testing"

	"github.com/caffeine-addictt/template/cmd/utils/types"
)

func TestNoParsing(t *testing.T) {
	val := "test"
	typeString := "my type"

	v := types.NewValueGuardNoParsing(val, typeString)
	if err := checkValues(val, typeString, v); err != nil {
		t.Fatal(err)
	}

	if err := v.Set("new value"); err != nil {
		t.Fatalf("failed to set value: %v", err)
	}
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

	if err := checkValues(val, typeString, v); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Only error when setting
	if err := v.Set(val); err == nil {
		t.Fatal("expected error, got nil")
	}
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

func checkValues(val, typeString string, vg *types.ValueGuard[string]) error {
	if vg.Value() != val {
		return fmt.Errorf("expected %s, got %s", val, vg.Value())
	}

	if vg.String() != val {
		return fmt.Errorf("expected %s, got %s", val, vg.String())
	}

	if vg.Type() != typeString {
		return fmt.Errorf("expected %s, got %s", typeString, vg.Type())
	}

	return nil
}
