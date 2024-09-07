package global_test

import (
	"regexp"
	"testing"

	"github.com/caffeine-addictt/waku/cmd/global"
	"github.com/stretchr/testify/assert"
)

// Regex taken from https://semver.org
var semverRegex = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

func TestFollowsSemVer(t *testing.T) {
	assert.True(t, semverRegex.MatchString(global.Version), "%s does not follow semver", global.Version)
}
