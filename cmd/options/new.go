package options

import (
	"errors"
	"os"
	"os/exec"

	"github.com/caffeine-addictt/template/cmd/utils/types"
)

// The options for the new command
var NewOpts = NewOptions{
	Repo:      *types.NewValueGuardNoParsing("https://github.com/caffeine-addictt/template", types.REPO),
	Branch:    *types.NewValueGuardNoParsing("", types.BRANCH),
	Directory: *types.NewValueGuardNoParsing("template", types.PATH),
}

type NewOptions struct {
	// The repository Url to use
	// Should be this repository by default
	Repo types.ValueGuard[string]

	// The branch to use
	Branch types.ValueGuard[string]

	// The directory of the template to use
	Directory types.ValueGuard[string]
}

// To clone the repository
func (o *NewOptions) CloneRepo(out chan string, e chan error) {
	tmpDirPath, err := os.MkdirTemp("", "template-*")
	if err != nil {
		e <- err
		return
	}

	var c *exec.Cmd
	if o.Branch.Value() != "" {
		c = exec.Command("git", "clone", "--depth", "1", "--branch", o.Branch.Value(), o.Repo.Value(), tmpDirPath)
	} else {
		c = exec.Command("git", "clone", "--depth", "1", o.Repo.Value(), tmpDirPath)
	}

	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		if errCleanup := os.RemoveAll(tmpDirPath); errCleanup != nil {
			e <- errors.Join(errCleanup, err)
		} else {
			e <- err
		}
		return
	}

	out <- tmpDirPath
	e <- nil
}
