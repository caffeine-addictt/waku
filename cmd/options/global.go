package options

// The global options for the CLI
var GlobalOpts = GlobalOptions{
	Verbosity:  0,
	Quiet:      false,
	Accessible: false,
}

type GlobalOptions struct {
	// The verbosity level of the CLI
	// The larger, the more verbose
	Verbosity int

	// Silence output
	Quiet bool

	// Mainly for screen-reader support, dropping
	// TUIs in favor of traditional prompts
	Accessible bool
}
