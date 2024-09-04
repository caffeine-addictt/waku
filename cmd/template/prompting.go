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

// ReadProjectName prompts user to enter project name
// or returns the name from options if it's provided.
func PromptForProjectName(name *string, projectRootDir *string) error {
	var state string
	if options.NewOpts.Name.Value() != "" {
		if err := validateProjectName(options.NewOpts.Name.Value(), name, projectRootDir, &state); err == nil {
			return nil
		}
	}

	return huh.NewForm(huh.NewGroup(
		huh.NewInput().TitleFunc(func() string {
			if state == "" {
				return "Name of your project"
			}

			return fmt.Sprintf("name of your project (%s)", state)
		}, &state).Validate(func(s string) error {
			err := validateProjectName(s, name, projectRootDir, &state)
			state = ""
			return err
		}),
	)).WithAccessible(options.GlobalOpts.Accessible).Run()
}

func validateProjectName(s string, name, projectRootDir, state *string) error {
	s = strings.TrimSpace(s)
	if s == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	*state = "resolving project location..."
	pDir, err := filepath.Rel(".", strings.ToLower(strings.ReplaceAll(s, " ", "-")))
	if err != nil {
		return err
	}

	*state = fmt.Sprintf("checking if directory '%s' already exists...", s)
	if ok, err := utils.PathExists(pDir); err != nil {
		return err
	} else if ok {
		return fmt.Errorf("directory '%s' already exists", s)
	}

	*projectRootDir = pDir
	*name = s
	return nil
}
