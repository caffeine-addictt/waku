package options

import (
	"errors"
	"os"

	"github.com/caffeine-addictt/waku/internal/git"
	"github.com/caffeine-addictt/waku/internal/log"
	"github.com/caffeine-addictt/waku/internal/types"
	"github.com/caffeine-addictt/waku/internal/utils"
	"github.com/caffeine-addictt/waku/pkg/version"
)

const defaultRepo = "https://github.com/caffeine-addictt/waku.git"

// The options for the new command
var NewOpts = NewOptions{
	Repo:      *types.NewValueGuard("", cmdOpt, types.REPO),
	Branch:    *types.NewValueGuard("", cmdOpt, types.BRANCH),
	Directory: *types.NewValueGuard("", cmdOpt, types.PATH),
	Name:      *types.NewValueGuard("", cmdOpt, types.STRING),
	License:   *types.NewValueGuard("", cmdOpt, types.STRING),
	Style:     *types.NewValueGuard("", cmdOpt, types.STRING),
	NoGit:     false,
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

	// Whether to skip initializing git
	NoGit bool
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
		"https://github.com/caffeine-addictt/waku",
		"git://github.com/caffeine-addictt/waku.git",
		"git@github.com:caffeine-addictt/waku.git":
		if err := o.Directory.Set("template"); err != nil {
			return err
		}
		if o.Branch.Value() == "" {
			if err := o.Branch.Set("v" + version.Version); err != nil {
				return err
			}
		}
	}

	return nil
}

// To clone the repository
func (o *NewOptions) CloneRepo() (string, error) {
	log.Debugln("creating tmp dir")

	tmpDirPath, err := os.MkdirTemp("", "template-*")
	if err != nil {
		return "", err
	}

	log.Infoln("create tmp dir at", tmpDirPath)

	opts := git.CloneOptions{
		Depth:     1,
		Branch:    o.Branch.Value(),
		Url:       utils.EscapeTermString(o.Repo.Value()),
		ClonePath: utils.EscapeTermString(tmpDirPath),
	}

	if err := git.Clone(opts); err != nil {
		if errCleanup := os.RemoveAll(tmpDirPath); errCleanup != nil {
			return "", errors.Join(errCleanup, err)
		}
		return "", err
	}

	return tmpDirPath, nil
}
