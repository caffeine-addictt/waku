package options

import (
	"fmt"

	"github.com/caffeine-addictt/template/cmd/utils"
	"github.com/caffeine-addictt/template/cmd/utils/types"
)

// The global options for the CLI
var GlobalOpts = GlobalOptions{
	Debug:   false,
	Verbose: false,
	Repo:    *types.NewValueGuardNoParsing("", "<repo>"),
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

type GlobalOptions struct {
	// The repository Url to use
	// Should be this repository by default
	Repo types.ValueGuard[string]

	// Where the cached repositories will live
	CacheDir types.ValueGuard[string]

	// Wheter or not debug mode should be enabled
	Debug bool

	// Wheter or not verbose mode should be enabled
	Verbose bool
}

// To resolve the options after the user has provided them
func (o *GlobalOptions) ResolveOptions() error {
	if err := o.resolveCacheDir(); err != nil {
		return err
	}

	return nil
}

func (o *GlobalOptions) resolveCacheDir() error {
	if o.CacheDir.Value() != "" {
		return nil
	}

	defaultCacheDir, err := utils.GetDefaultCacheDir()
	if err != nil {
		return err
	}

	return o.CacheDir.Set(defaultCacheDir)
}
