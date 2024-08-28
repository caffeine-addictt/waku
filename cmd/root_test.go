package cmd_test

import (
	"testing"

	"github.com/caffeine-addictt/template/cmd"
	"github.com/caffeine-addictt/template/cmd/helpers"
)

func TestRootCommandCanRun(t *testing.T) {
	stdout, stderr, err := helpers.ExecuteCommand(cmd.RootCmd, []string{})
	if err != nil {
		t.Fatalf("failed to run root command: %v", err)
	}

	if stdout == "" {
		t.Fatalf("expected non-empty stdout, but got: %s", stdout)
	}

	if stderr != "" {
		t.Fatalf("expected empty stderr, but got: %s", stderr)
	}
}
