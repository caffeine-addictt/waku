package git

import (
	"os"
	"os/exec"

	"github.com/caffeine-addictt/waku/internal/log"
	"github.com/go-git/go-git/v5"
)

// The main entrypoint for initializing the git repo
func Init(dir string) error {
	hasGit, err := HasGit()
	if err != nil {
		return err
	}

	log.Debugln("running git init")
	if hasGit {
		return initWithSystemGit(dir)
	}

	return initWithGoGet(dir)
}

func initWithSystemGit(dir string) error {
	c := exec.Command("git", "init")
	c.Stdin = os.Stdin
	c.Dir = dir

	if log.GetLevel() <= log.INFO {
		c.Stdout = os.Stdout
		c.Stdin = os.Stdin
	}

	return c.Run()
}

func initWithGoGet(dir string) error {
	_, err := git.PlainInit(dir, false)
	return err
}
