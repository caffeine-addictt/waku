package options

import (

	"github.com/caffeine-addictt/template/cmd/utils"
	"github.com/caffeine-addictt/template/cmd/utils/types"
)

// The options for the new command
var NewOpts = NewOptions{
	Repo:      *types.NewValueGuardNoParsing("https://github.com/caffeine-addictt/template", types.REPO),
	Branch:    *types.NewValueGuardNoParsing("", types.BRANCH),
	Directory: *types.NewValueGuardNoParsing("template", types.PATH),
}

type NewOptions struct {
	// The repository Url to use
	// Should be this repository by default
	Repo types.ValueGuard[string]

	// The branch to use
	Branch types.ValueGuard[string]

	// The directory of the template to use
	Directory types.ValueGuard[string]
}



	}

		return err
	}

}
