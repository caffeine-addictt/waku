package commands

import (
	"github.com/caffeine-addictt/template/cmd/global"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"ver"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(global.Version)
	},
}
