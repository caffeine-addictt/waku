package cmd_test

import (
	"testing"

	"github.com/caffeine-addictt/template/cmd"
	"github.com/caffeine-addictt/template/cmd/helpers"
	"github.com/stretchr/testify/assert"
)

func TestRootCommandCanRun(t *testing.T) {
	stdout, stderr, err := helpers.ExecuteCommand(cmd.RootCmd, []string{})

	assert.NoError(t, err, "failed to run root command")
	assert.NotEmpty(t, stdout, "expected empty stdout")
	assert.Empty(t, stderr, "expected empty stderr")
}
