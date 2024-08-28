package helpers_test

import (
	"testing"

	"github.com/caffeine-addictt/template/cmd/helpers"
	"github.com/spf13/cobra"
)

func TestExecuteCommandCapturesStderr(t *testing.T) {
	msg := "I'm in stderr"
	dummyCmd := cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			cmd.PrintErr(msg)
		},
	}

	stdout, stderr, err := helpers.ExecuteCommand(&dummyCmd, []string{}, "")
	if err != nil {
		t.Fatalf("failed to execute command: %v", stderr)
	}

	if stderr != msg {
		t.Fatalf("expected stderr to be '%s', got '%s'", msg, stderr)
	}

	if stdout != "" {
		t.Fatalf("expected stdout to be empty, got '%s'", stdout)
	}
}

func TestExecuteCommandCapturesStdout(t *testing.T) {
	msg := "I'm in stdout"
	dummyCmd := cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Print(msg)
		},
	}

	stdout, stderr, err := helpers.ExecuteCommand(&dummyCmd, []string{}, "")
	if err != nil {
		t.Fatalf("failed to execute command: %v", stderr)
	}

	if stdout != msg {
		t.Fatalf("expected stdout to be '%s', got '%s'", msg, stdout)
	}

	if stderr != "" {
		t.Fatalf("expected stderr to be empty, got '%s'", stderr)
	}
}
