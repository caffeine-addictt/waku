// Waku's internal wrapper for git
// aimed at supporting no-git sysstems
package git

import (
	"errors"
	"os/exec"

	e "github.com/caffeine-addictt/waku/internal/errors"
	"github.com/caffeine-addictt/waku/internal/log"
)

// HasGit checks if git is in PATH
func HasGit() (bool, error) {
	log.Infoln("checking if git is in PATH...")
	err := exec.Command("git", "--version").Run()
	if err != nil {
		if !errors.Is(err, exec.ErrNotFound) {
			return false, e.NewWakuErrorf("failed to check if git is in PATH: %v", err)
		}

		log.Debugln("git not found in $PATH, falling back to go-git")
		return false, nil
	}

	log.Debugln("found git in $PATH")
	return true, nil
}
