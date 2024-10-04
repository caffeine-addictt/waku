package cmd

import (
	"fmt"
	"os"

	"github.com/caffeine-addictt/waku/internal/errors"
	mango "github.com/muesli/mango-cobra"
	"github.com/muesli/roff"
	"github.com/spf13/cobra"
)

var ManCmd = &cobra.Command{
	Use:                   "man",
	Short:                 "generate Waku's manpages",
	Long:                  "Generate Waku's manpages.",
	SilenceUsage:          true,
	SilenceErrors:         true,
	DisableFlagsInUseLine: true,
	Hidden:                true,
	Args:                  cobra.NoArgs,
	ValidArgsFunction:     cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		manPage, err := mango.NewManPage(1, RootCmd)
		if err != nil {
			return errors.ToWakuError(err)
		}

		if _, err := fmt.Fprint(os.Stdout, manPage.Build(roff.NewDocument())); err != nil {
			return errors.ToWakuError(err)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(ManCmd)
}
