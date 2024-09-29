# Prompts

Prompts are additionally `asks` for templating.

!!! TIP

    Waku uses Go's [`text/template`](https://pkg.go.dev/text/template)
    library for templating, with 3 brackets `{{{ }}}` instead of 2.

    See the [documentation](https://pkg.go.dev/text/template#hdr-Text_and_spaces)
    for how you can leverage Go's strong templating engine.

## template.json

```json
{
  "prompts": [
    {
      // Key is case-sensitive and templated to.
      //
      // For example, using `key="MyKey"` will
      // match `{{{ .MyKey }}}` in files.
      //
      // (required)
      "key": "",

      // Type is the type of value key takes.
      // This can be either `str` or `arr`.
      //
      // If the type is `arr`, string input from the
      // user is split based on `sep`, allowing you
      // to do iteration with `{{{ range .MyKey }}}{{{ end }}}`.
      //
      // (required)
      "type": "",

      // Ask is the prompt that will be shown to the user.
      // (optional)
      "ask": "",

      // Sep is the separator used when type is `arr`.
      //
      // By default, Waku separates on space.
      //
      // (optional)
      "sep": " ",

      // Validate is the Regex used to validate the user input.
      //
      // By default, Waku checks for no empty input.
      //
      // (optional)
      "validate": ".+",

      // Capture is the Regex capture group used
      // to extract the value from the user input.
      // In the case of multiple captures, Waku only
      // uses the last capture.
      //
      // By default, Waku captures everything except
      // leading and trailing whitespace.
      //
      // (optional)
      "capture": "\s*(.*?)\s*",

      // Format is how each individual value(s)
      // are templated.
      //
      // All instances of `*` not directly after by a
      // backslash `\` are replaced with the capture value.
      //
      // For example, this is how you would ask for temperature
      // and ensure that the final templated value has the °C unit:
      //   capture: "\s*(-?\d+(?:\.\d+)?)\s*°?C?\s*"
      //   validate: "\s*(-?\d+(?:\.\d+)?)\s*°?C?\s*"
      //   format: "*°C"
      //
      // By default, Waku uses the capture value as is.
      //
      // (optional)
      "format": "*"
    }
  ]
}
```
