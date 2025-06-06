package cmd

import (
	"fmt"
	"os"

	"github.com/caffeine-addictt/waku/cmd/cleanup"
	"github.com/caffeine-addictt/waku/cmd/commands"
	"github.com/caffeine-addictt/waku/cmd/options"
	"github.com/caffeine-addictt/waku/internal/errors"
	"github.com/caffeine-addictt/waku/internal/utils"
	"github.com/caffeine-addictt/waku/pkg/log"
	"github.com/spf13/cobra"
)

// The Root command
var RootCmd = &cobra.Command{
	Use:           "waku",
	Short:         "let's make starting new projects feel like a breeze again",
	SilenceErrors: true,
	SilenceUsage:  true,
	Long: utils.MultilineString(
		"Waku (waku!) 枠組み",
		"",
		"Waku helps you kickstart new projects from templates.",
		"Let's make starting new projects feel like a breeze again.",
	),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if options.GlobalOpts.Quiet {
			return log.SetLevel(log.QUIET)
		}

		if options.GlobalOpts.Verbosity > 3 || options.GlobalOpts.Verbosity < 0 {
			return fmt.Errorf("verbosity level must be between 0 and 3")
		}

		return log.SetLevel(log.WARNING - log.Level(options.GlobalOpts.Verbosity))
	},
}

// Setting up configuration
func init() {
	RootCmd.PersistentFlags().BoolVarP(&options.GlobalOpts.Quiet, "quiet", "q", false, "quiet mode")
	RootCmd.PersistentFlags().BoolVarP(&options.GlobalOpts.Accessible, "accessible", "A", false, "accessible mode")
	RootCmd.PersistentFlags().CountVarP(&options.GlobalOpts.Verbosity, "verbose", "v", "verobisty level (v: info, vv: debug, vvv: trace)")

	RootCmd.MarkFlagsMutuallyExclusive("quiet", "verbose")

	commands.InitCommands(RootCmd)
}

// The main entry point for the command line tool
func Execute() {
	cleanup.On()

	var err error
	func() {
		defer func() {
			if r := recover(); r != nil {
				cleanup.Cleanup()
				cleanup.CleanupError()

				_ = log.SetLevel(log.TRACE) // force stack
				log.Fatalf("%v\n", r)
			}
		}()

		err = RootCmd.Execute()
	}()

	cleanup.Cleanup()
	if err != nil {
		cleanup.CleanupError()

		if _, ok := errors.IsWakuError(err); ok {
			log.Fatalf("%v\n", err)
		}

		cmd, _, err := RootCmd.Find(os.Args[1:])
		if err != nil {
			log.Fatalf("%v\n", err)
		} else {
			log.Fatalln(cmd.UsageString())
		}
	}
}
