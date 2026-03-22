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

| Usage                      | Description                                                                              |
| -------------------------- | ---------------------------------------------------------------------------------------- |
| `toLower "ME"`             | makes input string lowercase. See [ToLower](https://pkg.go.dev/strings#ToLower)          |
| `toUpper "me"`             | makes input string uppercase. See [ToUpper](https://pkg.go.dev/strings#ToUpper)          |
| `toTitle "my message"`     | makes input string titlecase. See [ToTitle](https://pkg.go.dev/strings#ToTitle)          |
| `trim "  hi  "`            | trims leading/trailing whitespace. See [TrimSpace](https://pkg.go.dev/strings#TrimSpace) |
| `replace "a-b" "-" "_"`    | replaces all occurrences. See [ReplaceAll](https://pkg.go.dev/strings#ReplaceAll)        |
| `contains "hello" "ell"`   | checks substring presence. See [Contains](https://pkg.go.dev/strings#Contains)           |
| `hasPrefix "hello" "he"`   | checks prefix. See [HasPrefix](https://pkg.go.dev/strings#HasPrefix)                     |
| `hasSuffix "hello" "lo"`   | checks suffix. See [HasSuffix](https://pkg.go.dev/strings#HasSuffix)                     |
| `join (slice "a" "b") ","` | joins slice into string. See [Join](https://pkg.go.dev/strings#Join)                     |
| `split "a,b" ","`          | splits string into slice. See [Split](https://pkg.go.dev/strings#Split)                  |
| `slug "Hello World"`       | lowercase + replaces spaces with `-`                                                     |

### Arithmetic

_All inputs are strings, output is string._

| Usage          | Description                   |
| -------------- | ----------------------------- |
| `add "1" "2"`  | addition                      |
| `sub "5" "3"`  | subtraction                   |
| `mul "2" "4"`  | multiplication                |
| `div "10" "2"` | division (IEEE-754, no panic) |

### Logic / Flow

| Usage                   | Description                      |
| ----------------------- | -------------------------------- |
| `ternary cond a b`      | returns `a` if true, else `b`    |
| `default "" "fallback"` | returns fallback if empty string |

### Time

| Usage                        | Description                                                                 |
| ---------------------------- | --------------------------------------------------------------------------- |
| `timefmt .Time "2006-01-02"` | formats `time.Time`. See [Time.Format](https://pkg.go.dev/time#Time.Format) |

### Encoding

| Usage    | Description                                                                             |
| -------- | --------------------------------------------------------------------------------------- |
| `json .` | pretty JSON encode. See [MarshalIndent](https://pkg.go.dev/encoding/json#MarshalIndent) |
