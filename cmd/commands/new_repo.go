package commands

import "github.com/spf13/cobra"

var NewRepoCmd = &cobra.Command{
	Use:     "repo",
	Aliases: []string{"repository", "project"},
	Short:   NewCmd.Short,
	Long:    NewCmd.Long,
	Run:     NewCmd.Run,
}

func init() {
	AddNewCmdFlags(NewRepoCmd)
	NewCmd.AddCommand(NewRepoCmd)
}
