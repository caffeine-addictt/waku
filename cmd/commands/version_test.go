package commands_test

import (
	"testing"

	"github.com/caffeine-addictt/template/cmd/commands"
	"github.com/caffeine-addictt/template/cmd/global"
	"github.com/caffeine-addictt/template/cmd/helpers"
)

func TestVersionOut(t *testing.T) {
	stdout, stderr, err := helpers.ExecuteCommand(commands.VersionCmd, []string{})
	if err != nil {
		t.Fatal(err)
	}

	if stdout != global.Version+"\n" {
		t.Errorf("Expected version %s, got %s", global.Version, stdout)
	}

	if stderr != "" {
		t.Fatalf("Expected no stderr, got %s", err)
	}
}
