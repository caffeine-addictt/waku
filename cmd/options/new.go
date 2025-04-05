package options

import (
	"errors"
	"os"

	"github.com/caffeine-addictt/waku/cmd/cleanup"
	e "github.com/caffeine-addictt/waku/internal/errors"
	"github.com/caffeine-addictt/waku/internal/git"
	"github.com/caffeine-addictt/waku/internal/types"
	"github.com/caffeine-addictt/waku/internal/utils"
	"github.com/caffeine-addictt/waku/pkg/log"
	"github.com/caffeine-addictt/waku/pkg/version"
)

const defaultRepo = "https://github.com/caffeine-addictt/waku.git"

// The options for the new command
var NewOpts = NewOptions{
	Repo:        *types.NewValueGuard("", cmdOpt, types.REPO),
	Source:      *types.NewValueGuard("", cmdOpt, types.REPO),
	Branch:      *types.NewValueGuard("", cmdOpt, types.BRANCH),
	Directory:   *types.NewValueGuard("", cmdOpt, types.PATH),
	Name:        *types.NewValueGuard("", cmdOpt, types.STRING),
	License:     *types.NewValueGuard("", cmdOpt, types.STRING),
	Style:       *types.NewValueGuard("", cmdOpt, types.STRING),
	NoGit:       false,
	NoLicense:   false,
	AllowSpaces: false,
}

type NewOptions struct {
	// The repository Url to use
	// Should be this repository by default
	Repo types.ValueGuard[string]

	// The repository Url or local path to use
	// Should be this repository by default
	Source types.ValueGuard[string]

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

	// Whether to skip initializing license
	NoLicense bool

	// Whether to allow spaces in the project name
	// This will skip replacing them with hyphens
	AllowSpaces bool
}

func cmdOpt(v string) (string, error) {
	return utils.CleanStringStrict(v), nil
}

// TO be invoked before a command is ran
func (o *NewOptions) Validate() error {
	// Since both flags are mutually exclusive
	if err := o.Source.Set(o.Source.Value() + o.Repo.Value()); err != nil {
		return err
	}

	switch o.Source.Value() {
	case "":
		if err := o.Source.Set(defaultRepo); err != nil {
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
			newBranch := "v" + version.Version
			log.Debugf("Setting branch to %s\n", newBranch)
			if err := o.Branch.Set(newBranch); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetSource returns the source directory path
// that is either cloned with Git or is local.
//
// If it is a Git cloned path, it will be cleaned
func (o *NewOptions) GetSource() (string, error) {
	switch git.CheckUrl(o.Source.Value()) {
	case git.GitUrlType:
		s, err := o.CloneRepo()
		if err != nil {
			return s, e.NewWakuErrorf("could not clone repo: %v", err)
		}

		cleanup.Schedule(func() error {
			log.Debugf("removing tmp dir: %s\n", s)
			if err := os.RemoveAll(s); err != nil {
				return e.NewWakuErrorf("failed to cleanup tmp dir: %v", err)
			}
			return nil
		})

		return s, err

	case git.PathUrlType:
		return o.Source.Value(), nil
	}

	return "", e.NewWakuErrorf("invalid source URL or path: %s", o.Source.Value())
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
		Url:       utils.EscapeTermString(o.Source.Value()),
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
