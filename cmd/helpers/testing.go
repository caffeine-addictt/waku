package helpers

import (
	"bytes"
	"strings"

	"github.com/spf13/cobra"
)

// For testing command execution
// Returns stdout, stderr and error
func ExecuteCommand(cmd *cobra.Command, stdin []string, args ...string) (string, string, error) {
	cmd.SetArgs(args)

	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}

	cmd.SetOut(&stdout)
	cmd.SetErr(&stderr)
	cmd.SetIn(strings.NewReader(strings.Join(stdin, "\n")))

	err := cmd.Execute()
	if err != nil {
		return stdout.String(), stderr.String(), err
	}

	return stdout.String(), stderr.String(), nil
}
