package commands

import "github.com/spf13/cobra"

var NewRepoCmd = &cobra.Command{
	Use:           "repo",
	Aliases:       []string{"repository", "project"},
	SilenceErrors: true,
	SilenceUsage:  true,
	Short:         NewCmd.Short,
	Long:          NewCmd.Long,
	PreRunE:       NewCmd.PreRunE,
	Run:           NewCmd.Run,
}

func init() {
	AddNewCmdFlags(NewRepoCmd)
	NewCmd.AddCommand(NewRepoCmd)
}
