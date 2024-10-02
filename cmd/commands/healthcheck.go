package commands

import (
	"errors"
	"os/exec"

	e "github.com/caffeine-addictt/waku/internal/errors"
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
		if err := exec.Command("git", "--version").Run(); err != nil {
			if !errors.Is(err, exec.ErrNotFound) {
				return e.NewWakuErrorf(err.Error())
			}

			log.Warnln("Git not found in $PATH, Waku will fallback to go-git. Authentication may not work as expected.")
		}

		return nil
	},
}
