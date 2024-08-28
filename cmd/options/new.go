package options

import (
	"fmt"

	"github.com/caffeine-addictt/template/cmd/utils"
	"github.com/caffeine-addictt/template/cmd/utils/types"
)

// The options for the new command
var NewOpts = NewOptions{
	Repo:      *types.NewValueGuardNoParsing("", types.REPO),
	Branch:    *types.NewValueGuardNoParsing("", types.BRANCH),
	Directory: *types.NewValueGuardNoParsing("", types.PATH),
	CacheDir: *types.NewValueGuard("", func(v string) (string, error) {
		ok, err := utils.IsDir(v)
		if err != nil {
			return "", err
		}

		if !ok {
			return "", fmt.Errorf("'%s' is not a valid directory", v)
		}

		return v, nil
	}, types.PATH),
}

type NewOptions struct {
	// The repository Url to use
	// Should be this repository by default
	Repo types.ValueGuard[string]

	// The branch to use
	Branch types.ValueGuard[string]

	// The directory of the template to use
	Directory types.ValueGuard[string]

	// Where the cached repositories will live
	CacheDir types.ValueGuard[string]
}

// To resolve the options after the user has provided them
func (o *NewOptions) ResolveOptions() error {
	if err := o.resolveCacheDir(); err != nil {
		return err
	}

	return nil
}

func (o *NewOptions) resolveCacheDir() error {
	if o.CacheDir.Value() != "" {
		return nil
	}

	defaultCacheDir, err := utils.GetDefaultCacheDir()
	if err != nil {
		return err
	}

	return o.CacheDir.Set(defaultCacheDir)
}
