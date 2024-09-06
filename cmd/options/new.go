package options

import (
	"errors"
	"os"
	"os/exec"

	"github.com/caffeine-addictt/template/cmd/utils"
	"github.com/caffeine-addictt/template/cmd/utils/types"
)

const defaultRepo = "https://github.com/caffeine-addictt/template.git"

// The options for the new command
var NewOpts = NewOptions{
	Repo:      *types.NewValueGuard("", cmdOpt, types.REPO),
	Branch:    *types.NewValueGuard("", cmdOpt, types.BRANCH),
	Directory: *types.NewValueGuard("", cmdOpt, types.PATH),
	Name:      *types.NewValueGuard("", cmdOpt, types.STRING),
	License:   *types.NewValueGuard("", cmdOpt, types.STRING),
	Style:     *types.NewValueGuard("", cmdOpt, types.STRING),
}

type NewOptions struct {
	// The repository Url to use
	// Should be this repository by default
	Repo types.ValueGuard[string]

	// The branch to use
	Branch types.ValueGuard[string]

	// The directory of the template to use
	Directory types.ValueGuard[string]

	// The name of your project
	Name types.ValueGuard[string]

	// The license of your project
	License types.ValueGuard[string]

	// The style to use
	Style types.ValueGuard[string]
}

func cmdOpt(v string) (string, error) {
	return utils.CleanString(v), nil
}

// TO be invoked before a command is ran
func (o *NewOptions) Validate() error {
	switch o.Repo.Value() {
	case "":
		if err := o.Repo.Set(defaultRepo); err != nil {
			return err
		}
		if err := o.Directory.Set("template"); err != nil {
			return err
		}

	case
		defaultRepo,
		"https://github.com/caffeine-addictt/template",
		"git://github.com/caffeine-addictt/template.git",
		"git@github.com:caffeine-addictt/template.git":
		if err := o.Directory.Set("template"); err != nil {
			return err
		}
	}

	return nil
}

// To clone the repository
func (o *NewOptions) CloneRepo() (string, error) {
	Debugln("Creating tmp dir")

	tmpDirPath, err := os.MkdirTemp("", "template-*")
	if err != nil {
		return "", err
	}

	Infoln("Create tmp dir at", tmpDirPath)

	args := []string{"clone", "--depth", "1"}
	if o.Branch.Value() != "" {
		args = append(args, "--branch", utils.EscapeTermString(o.Branch.Value()))
	}
	args = append(args, utils.EscapeTermString(o.Repo.Value()), utils.EscapeTermString(tmpDirPath))

	Debugln("git args:", args, len(args))

	c := exec.Command("git", args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		if errCleanup := os.RemoveAll(tmpDirPath); errCleanup != nil {
			return "", errors.Join(errCleanup, err)
		}
		return "", err
	}

	return tmpDirPath, nil
}
