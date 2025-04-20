# {{{ .Name }}}

{{{ .Description }}}

## Getting Started

{{{ .GettingStarted }}}

## Features

{{{- if gt (len .Features) 0 }}}
{{{- range .Features }}}

- {{{ . }}}
{{{- end }}}
{{{- end }}}

{{{- if .License }}}

## License

{{{ .License }}}
{{{ end }}}
