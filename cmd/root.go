package cmd

import (
	"log"

	"github.com/caffeine-addictt/waku/cmd/commands"
	"github.com/caffeine-addictt/waku/cmd/options"
	"github.com/caffeine-addictt/waku/cmd/utils"
	"github.com/spf13/cobra"
)

// The Root command
var RootCmd = &cobra.Command{
	Use:   "waku",
	Short: "let's make starting new projects feel like a breeze again",
	Long: utils.MultilineString(
		"Waku (waku!) 枠組み",
		"",
		"Waku helps you kickstart new projects from templates.",
		"Let's make starting new projects feel like a breeze again.",
	),
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
