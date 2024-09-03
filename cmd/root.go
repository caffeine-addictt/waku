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
	Short: "let's make starting new projects feel like a breeze again",
	Long:  "This tool helps you to create a new project from templates.\n\nLet's make starting new projects feel like a breeze again.",
}

// Setting up configuration
func init() {
	RootCmd.PersistentFlags().BoolVarP(&options.GlobalOpts.Debug, "debug", "d", false, "debug mode")
	RootCmd.PersistentFlags().BoolVarP(&options.GlobalOpts.Verbose, "verbose", "v", false, "verbose mode")
	RootCmd.PersistentFlags().BoolVarP(&options.GlobalOpts.Accessible, "accessible", "A", false, "accessible mode")

	commands.InitCommands(RootCmd)
}

// The main entry point for the command line tool
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
