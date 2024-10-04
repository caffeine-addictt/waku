package commands

import (
	"path/filepath"
	"strings"

	"github.com/caffeine-addictt/waku/internal/errors"
	"github.com/caffeine-addictt/waku/internal/template"
	"github.com/caffeine-addictt/waku/internal/utils"
	"github.com/spf13/cobra"
)

var CheckCmd = &cobra.Command{
	Use:           "check <path>",
	Aliases:       []string{"ch", "c", "verify"},
	Short:         "check if template.json is valid",
	Long:          "Check if your current template.json is valid",
	Args:          cobra.MaximumNArgs(1),
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check for naming
		if len(args) == 1 && !strings.HasSuffix(args[0], "template.json") {
			return errors.NewWakuErrorf("name your file template.json")
		}

		// Resolve file path
		var filePath string
		if len(args) == 1 {
			filePath = args[0]
		} else {
			filePath = "template.json"
		}
		filePath = filepath.Clean(filePath)

		ok, err := utils.IsFile(filePath)
		if err != nil {
			return errors.NewWakuErrorf("failed to check if %s is a file: %v", filePath, err)
		}
		if !ok {
			return errors.NewWakuErrorf("%s does not exist or is not a file", filePath)
		}

		if _, err := template.ParseConfig(filePath); err != nil {
			return errors.ToWakuError(err)
		}

		cmd.Println("Seems ok!")
		return nil
	},
}
