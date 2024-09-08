package cmd

import (
	"os"
	"path/filepath"

	"github.com/caffeine-addictt/waku/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var ManCmd = &cobra.Command{
	Use:   "man",
	Short: "generate man pages",
	Long: utils.MultilineString(
		"Generate man pages for the command.",
		"By default, man pages are generated in /usr/share/man/man1.",
	),
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirPathDirty := "/usr/share/man/man1/"
		if len(args) == 1 {
			dirPathDirty = args[0]
		}

		dirPath := filepath.Clean(dirPathDirty)
		if err := doc.GenManTree(RootCmd, &doc.GenManHeader{
			Title:   "Waku Command Reference",
			Source:  "Waku (waku!) 枠組み",
			Section: "1",
		}, dirPath); err != nil {
			cmd.PrintErrf("could not create man pages: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(ManCmd)
}
