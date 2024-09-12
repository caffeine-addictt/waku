package cmd

import (
	"fmt"
	"os"

	mango "github.com/muesli/mango-cobra"
	"github.com/muesli/roff"
	"github.com/spf13/cobra"
)

var ManCmd = &cobra.Command{
	Use:                   "man",
	Short:                 "generate Waku's manpages",
	Long:                  "Generate Waku's manpages.",
	SilenceUsage:          true,
	DisableFlagsInUseLine: true,
	Hidden:                true,
	Args:                  cobra.NoArgs,
	ValidArgsFunction:     cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		manPage, err := mango.NewManPage(1, RootCmd)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, manPage.Build(roff.NewDocument()))
		return err
	},
}

func init() {
	RootCmd.AddCommand(ManCmd)
}
