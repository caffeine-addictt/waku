package options

import "github.com/caffeine-addictt/template/cmd/utils/types"

type Options struct {
	// The repository Url to use
	// Should be this repository by default
	Repo types.ValueGuard[string]

	// Wheter or not debug mode should be enabled
	Debug bool
}
