package cmd

import (
	"log"

	"github.com/caffeine-addictt/template/cmd/commands"
	"github.com/caffeine-addictt/template/cmd/options"
	"github.com/spf13/cobra"
)

// The Root command
var RootCmd = &cobra.Command{
	Use:   "template",
	Short: "Let's make starting new projects feel like a breeze again.",
	Long:  "This tool helps you to create a new project from templates.",
}

// Setting up configuration
func init() {
	RootCmd.PersistentFlags().BoolVarP(&options.Opts.Debug, "debug", "d", false, "Debug mode")
	RootCmd.PersistentFlags().VarP(&options.Opts.Repo, "repo", "r", "Community source repository for templates")

	commands.InitCommands(RootCmd)
}

// The main entry point for the command line tool
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
