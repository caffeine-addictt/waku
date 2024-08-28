package options

// The global options for the CLI
var GlobalOpts = GlobalOptions{
	Debug:   false,
	Verbose: false,
}

type GlobalOptions struct {
	// Wheter or not debug mode should be enabled
	Debug bool

	// Wheter or not verbose mode should be enabled
	Verbose bool
}
