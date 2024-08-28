package commands

import (
	"os"

	"github.com/caffeine-addictt/template/cmd/options"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var NewCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"init"},
	Short:   "create a new project",
	Long:    "Create a new project from a template",
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := options.NewOpts.ResolveOptions(); err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	NewCmd.PersistentFlags().VarP(&options.NewOpts.Repo, "repo", "r", "community source repository for templates")
	NewCmd.PersistentFlags().VarP(&options.NewOpts.CacheDir, "cache", "C", "where source repository will be cloned to [default: $XDG_CONFIG_HOME/template]")
}
