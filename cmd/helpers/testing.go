package helpers

import (
	"bytes"
	"strings"

	"github.com/caffeine-addictt/waku/pkg/log"
	"github.com/spf13/cobra"
)

// For testing command execution
func ExecuteCommand(cmd *cobra.Command, stdin []string, args ...string) (stdout, stderr string, e error) {
	cmd.SetArgs(args)

	out := bytes.Buffer{}
	errout := bytes.Buffer{}

	cmd.SetOut(&out)
	cmd.SetErr(&errout)
	cmd.SetIn(strings.NewReader(strings.Join(stdin, "\n")))
	log.Stdout = &out
	log.Stderr = &errout

	err := cmd.Execute()
	if err != nil {
		return out.String(), errout.String(), err
	}

	return out.String(), errout.String(), nil
}
