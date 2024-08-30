package commands

import "github.com/spf13/cobra"

// To initialize all the commands as subcommands of root
func InitCommands(root *cobra.Command) {
	root.AddCommand(VersionCmd)
	root.AddCommand(NewCmd)
	root.AddCommand(CheckCmd)
}
