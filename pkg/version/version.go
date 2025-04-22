package version

import (
	"runtime/debug"
	"strings"
)

// The current app version
var Version string

const defaultVersion = "dev"

func init() {
	Version = getVersion()
}

func main() {
	Version = getVersion()
}

func getVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "(devel)" && info.Main.Version != "" {

			p := strings.Split(info.Main.Version, "-")
			p = p[:len(p)-2]

			return strings.TrimPrefix(strings.Join(p, "-"), "v")
		}
	}

	return defaultVersion
}

func GetBranch() string {
	if Version == defaultVersion {
		return "main"
	}

	return "v" + Version
}
