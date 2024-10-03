package commands

import (
	"github.com/caffeine-addictt/waku/internal/git"
	"github.com/caffeine-addictt/waku/internal/log"
	"github.com/spf13/cobra"
)

var HealthcheckCmd = &cobra.Command{
	Use:           "healthcheck",
	Aliases:       []string{"health"},
	Args:          cobra.NoArgs,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ok, err := git.HasGit()
		if err != nil {
			return err
		}

		if !ok {
			log.Warnln("Git not found in $PATH, Waku will fallback to go-git. Authentication may not work as expected.")
		}

		return nil
	},
}
