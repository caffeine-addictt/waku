# Variables

Variables are reusable values injected when templating
without a user prompt.

They're useful for setting constants in your styles.

!!! TIP

    Waku provides some default template values for you.
    See the docs [here](../templates.md).

## waku config

=== "Yaml"

    ```yaml
    styles:
      my-style:
        variables:
          # Key is case-sensitive and templated to.
          #
          # For example, using `key="MyKey"` will
          # match `{{{ .Variables.MyKey }}}` in files.
          #
          # (required)
          - key: ""

          # Format is what is used to format values to {{{ .Variables.X }}}.
          #
          # Format supports go's text/template templating and exposes
          # all default values and 'styles.prompts'.
          #
          # (required)
            fmt: ""

          # Type is the type of value key takes.
          # This can be either `str` or `arr`.
          #
          # If the type is `arr`, string input from the
          # user is split based on `sep`, allowing you
          # to do iteration with `{{{ range .Variables.MyKey }}}{{{ end }}}`.
          #
          # (optional)
            type: "str"

          # Sep is the separator used when type is `arr`.
          #
          # By default, Waku separates on space.
          #
          # (optional)
            sep: " "
    ```

=== "Json"

    ```json
    {
      "prompts": [
        {
          // Key is case-sensitive and templated to.
          //
          // For example, using `key="MyKey"` will
          // match `{{{ .Variables.MyKey }}}` in files.
          //
          // (required)
          "key": "",

          // Format is what is used to format values to {{{ .Variables.X }}}.
          //
          // Format supports go's text/template templating and exposes
          // all default values and 'styles.prompts'.
          //
          // (required)
          "fmt": "",

          // Type is the type of value key takes.
          // This can be either `str` or `arr`.
          //
          // If the type is `arr`, string input from the
          // user is split based on `sep`, allowing you
          // to do iteration with `{{{ range .Variables.MyKey }}}{{{ end }}}`.
          //
          // (optional)
          "type": "str",

          // Sep is the separator used when type is `arr`.
          //
          // By default, Waku separates on space.
          //
          // (optional)
          "sep": " "
        }
      ]
    }
    ```
