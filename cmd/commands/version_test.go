package commands_test

import (
	"testing"

	"github.com/caffeine-addictt/template/cmd/commands"
	"github.com/caffeine-addictt/template/cmd/global"
	"github.com/caffeine-addictt/template/cmd/helpers"
	"github.com/stretchr/testify/assert"
)

func TestVersionOut(t *testing.T) {
	stdout, stderr, err := helpers.ExecuteCommand(commands.VersionCmd, []string{})
	assert.NoError(t, err)

	assert.Equal(t, stdout, global.Version+"\n")
	assert.Empty(t, stderr)
}
