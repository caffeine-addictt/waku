package template

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/caffeine-addictt/template/cmd/options"
	"github.com/caffeine-addictt/template/cmd/utils"
	"github.com/charmbracelet/huh"
)

func ProptForLicense() error {
	return huh.NewForm().Run()
}

// PromptForProjectName prompts user to enter project name
// or returns the name from options if it's provided.
func PromptForProjectName(name *string, projectRootDir *string) *huh.Input {
	if options.NewOpts.Name.Value() != "" {
		if err := validateProjectName(options.NewOpts.Name.Value(), name, projectRootDir); err == nil {
			return nil
		}
	}

	return huh.NewInput().Title("Name of your project").Validate(func(s string) error {
		return validateProjectName(s, name, projectRootDir)
	})
}

func validateProjectName(s string, name, projectRootDir *string) error {
	s = strings.TrimSpace(s)
	if s == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	pDir, err := filepath.Rel(".", strings.ToLower(strings.ReplaceAll(s, " ", "-")))
	if err != nil {
		return err
	}

	if ok, err := utils.PathExists(pDir); err != nil {
		return err
	} else if ok {
		return fmt.Errorf("directory '%s' already exists", s)
	}

	*projectRootDir = pDir
	*name = s
	return nil
}
