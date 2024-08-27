package options

import (
	"github.com/caffeine-addictt/template/cmd/utils/types"
)

// The global options for the CLI
var Opts = Options{
	Debug: false,
	Repo:  *types.NewValueGuardNoParsing("", "<repo>"),
}
