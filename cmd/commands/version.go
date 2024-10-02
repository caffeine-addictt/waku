package commands

import (
	"github.com/caffeine-addictt/waku/cmd/global"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:           "version",
	Aliases:       []string{"ver"},
	Short:         "show version",
	Long:          "Show version of waku",
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(global.Version)
	},
}
