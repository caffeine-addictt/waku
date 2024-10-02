package commands

import (
	"os/exec"

	e "github.com/caffeine-addictt/waku/internal/errors"
	"github.com/spf13/cobra"
)

var HealthcheckCmd = &cobra.Command{
	Use:           "healthcheck",
	Aliases:       []string{"health"},
	Args:          cobra.NoArgs,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := exec.Command("git", "--version").Run(); err != nil {
			return e.NewWakuErrorf(err.Error())
		}

		return nil
	},
}
