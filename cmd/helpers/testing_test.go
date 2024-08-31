package helpers_test

import (
	"testing"

	"github.com/caffeine-addictt/template/cmd/helpers"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestExecuteCommandCapturesStderr(t *testing.T) {
	msg := "I'm in stderr"
	dummyCmd := cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			cmd.PrintErr(msg)
		},
	}

	stdout, stderr, err := helpers.ExecuteCommand(&dummyCmd, []string{}, "")
	assert.NoError(t, err, "failed to execute command")

	assert.Equal(t, msg, stderr, "wrong stderr")
	assert.Empty(t, stdout, "non-empty stdout")
}

func TestExecuteCommandCapturesStdout(t *testing.T) {
	msg := "I'm in stdout"
	dummyCmd := cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Print(msg)
		},
	}

	stdout, stderr, err := helpers.ExecuteCommand(&dummyCmd, []string{}, "")
	assert.NoError(t, err, "failed to execute command")

	assert.Equal(t, msg, stdout, "wrong stdout")
	assert.Empty(t, stderr, "non-empty stderr")
}
