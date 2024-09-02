package commands

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/caffeine-addictt/template/cmd/template"
	"github.com/caffeine-addictt/template/cmd/utils"
	"github.com/spf13/cobra"
)

var CheckCmd = &cobra.Command{
	Use:     "check <path>",
	Aliases: []string{"ch", "c", "verify"},
	Short:   "check if template.json is valid",
	Long:    "Check if your current template.json is valid",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Check for naming
		if len(args) == 1 && !strings.HasSuffix(args[0], "template.json") {
			cmd.PrintErrln("name your file template.json")
			os.Exit(1)
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
			cmd.PrintErrln(err)
			os.Exit(1)
		}
		if !ok {
			cmd.PrintErrln("template.json not found")
			os.Exit(1)
		}

		if _, err := template.ParseConfig(filePath); err != nil {
			cmd.PrintErrln(err)
			os.Exit(1)
		}

		cmd.Println("Seems ok!")
	},
}
