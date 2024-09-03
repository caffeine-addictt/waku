package options

// The global options for the CLI
var GlobalOpts = GlobalOptions{
	Debug:      false,
	Verbose:    false,
	Accessible: false,
}

type GlobalOptions struct {
	// Wheter or not debug mode should be enabled
	Debug bool

	// Wheter or not verbose mode should be enabled
	Verbose bool

	// Mainly for screen-reader support, dropping
	// TUIs in favor of traditional prompts
	Accessible bool
}

func (o *GlobalOptions) DebugOrVerbose() bool {
	return o.Debug || o.Verbose
}
