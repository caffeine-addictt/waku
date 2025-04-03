package commands

import (
	"github.com/caffeine-addictt/waku/internal/errors"
	"github.com/caffeine-addictt/waku/internal/template"
	"github.com/caffeine-addictt/waku/pkg/log"
	"github.com/spf13/cobra"
)

var CheckCmd = &cobra.Command{
	Use:           "check <path>",
	Aliases:       []string{"ch", "c", "verify"},
	Short:         "check if config is valid",
	Long:          "Check if your current config is valid",
	Args:          cobra.MaximumNArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var filePath string
		if len(args) == 1 {
			filePath = args[0]
		} else {
			filePath = "."
		}

		log.Debugf("checking if %s is a valid template\n", filePath)
		if _, _, err := template.ParseConfig(filePath); err != nil {
			return errors.ToWakuError(err)
		}

		log.Println("Seems ok!")
		return nil
	},
}
