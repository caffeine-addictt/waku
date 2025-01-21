package git

import (
	"net/url"
	"os"

	"github.com/caffeine-addictt/waku/pkg/log"
)

type UrlType string

const (
	GitUrlType  UrlType = "git"
	PathUrlType UrlType = "path"
	BadUrlType  UrlType = "bad"
)

var validSchemas [4]string = [4]string{"http", "https", "git", "ssh"}

// CheckUrl checks if a given string s is
// a valid git or path url
func CheckUrl(s string) UrlType {
	if IsGitUrl(s) {
		return GitUrlType
	}

	if IsPathUrl(s) {
		return PathUrlType
	}

	return BadUrlType
}

// IsPathUrl checks if a given string s is
// a valid fs path
func IsPathUrl(s string) bool {
	log.Debugf("checking if %s is a valid path\n", s)
	_, err := os.Stat(s)
	v := !os.IsNotExist(err)

	if !v {
		log.Debugf("%s is not a valid path", s)
	}

	return v
}

// IsGitUrl checks if a given string s is
// a valid Git url
func IsGitUrl(s string) bool {
	log.Debugf("checking if %s is a valid git url\n", s)
	parsedUrl, err := url.Parse(s)
	if err != nil {
		log.Debugf("%s is not a valid url\n", s)
		return false
	}

	for _, schema := range validSchemas {
		if parsedUrl.Scheme == schema {
			return true
		}
	}

	return false
}
