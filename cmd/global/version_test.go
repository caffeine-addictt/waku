package global_test

import (
	"regexp"
	"testing"

	"github.com/caffeine-addictt/template/cmd/global"
)

// Regex taken from https://semver.org
var semverRegex = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

func TestFollowsSemVer(t *testing.T) {
	if !semverRegex.MatchString(global.Version) {
		t.Fatalf("%v does not follow semver", global.Version)
	}
}
