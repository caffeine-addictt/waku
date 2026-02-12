# Templates

!!! TIP

    Waku uses Go's [`text/template`](https://pkg.go.dev/text/template)
    library for templating, with 3 brackets `{{{ }}}` instead of 2.

    See the [documentation](https://pkg.go.dev/text/template#hdr-Text_and_spaces)
    for how you can leverage Go's strong templating engine.

## Fields

You can use them in your styles by having `{{{ .Key }}}`
in your files.

| Key          | Description                                                         |
| ------------ | ------------------------------------------------------------------- |
| `.Name`      | the project name                                                    |
| `.License`   | the project license text (i.e. MIT License)                         |
| `.Spdx`      | the project license SPDX identifier (i.e. MIT)                      |
| `.Variables` | an object containing all defined [variables](./fields/variables.md) |

## Functions

We also supply our own custom functions in addition to Go's
[default functions](https://pkg.go.dev/text/template#hdr-Functions).

### String Operations

| Usage           | Description                                                         |
| ------------  | ------------------------------------------------------------ |
| `toLower "ME"`        | makes input string lowercase. See [ToLower](https://pkg.go.dev/strings#ToLower)                                                    |
| `toUpper "me"`        | makes input string uppercase. See [ToUpper](https://pkg.go.dev/strings#ToUpper)                                                    |
| `toTitle "my message"`        | makes input string titlecase. See [ToUpper](https://pkg.go.dev/golang.org/x/text/cases#Title)                                                    |
| `.License`    | the project license text (i.e. MIT License)                         |
| `.Spdx`       | the project license SPDX identifier (i.e. MIT)                      |
| `.Variables`   |an object containing all defined [variables](./fields/variables.md) |
