package git

import (
	"os"
	"os/exec"

	"github.com/caffeine-addictt/waku/internal/log"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type CloneOptions struct {
	Branch    string
	Url       string
	ClonePath string
	Depth     int
}

// The main entrypoint for cloning.
//
// Will use system Git if available, otherwise falls back to go-git
func Clone(opts CloneOptions) error {
	hasGit, err := HasGit()
	if err != nil {
		return err
	}

	log.Debugf("running git clone with options: %v\n", opts)
	if hasGit {
		return cloneWithSystemGit(opts)
	}

	return cloneWithGoGit(opts)
}

func cloneWithSystemGit(opts CloneOptions) error {
	args := []string{"clone", "--depth", "1"}

	if opts.Branch != "" {
		args = append(args, "--branch", opts.Branch)
	}
	args = append(args, opts.Url, opts.ClonePath)

	c := exec.Command("git", args...)
	c.Stdin = os.Stdin

	if log.GetLevel() <= log.INFO {
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
	}

	return c.Run()
}

func cloneWithGoGit(opts CloneOptions) error {
	args := git.CloneOptions{
		URL:   opts.Url,
		Depth: opts.Depth,
	}

	if opts.Branch != "" {
		args.ReferenceName = plumbing.ReferenceName(opts.Branch)
	}

	if log.GetLevel() <= log.INFO {
		args.Progress = os.Stdout
	}

	_, err := git.PlainClone(opts.ClonePath, false, &args)
	return err
}
