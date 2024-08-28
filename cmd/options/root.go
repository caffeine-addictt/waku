package options

import (
	"fmt"

	"github.com/caffeine-addictt/template/cmd/utils"
	"github.com/caffeine-addictt/template/cmd/utils/types"
)

// The global options for the CLI
var Opts = Options{
	Debug: false,
	Repo:  *types.NewValueGuardNoParsing("", "<repo>"),
	CacheDir: *types.NewValueGuard("", func(v string) (string, error) {
		ok, err := utils.IsDir(v)
		if err != nil {
			return "", err
		}

		if !ok {
			return "", fmt.Errorf("'%s' is not a valid directory", v)
		}

		return v, nil
	}, "<path>"),
}
