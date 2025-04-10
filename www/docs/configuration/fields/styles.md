# Styles

You can specify multiple styles by using the
`styles` field in your `waku config`.

The [ignore](./ignore.md), [includes](./includes.md), [labels](./labels.md), [prompts](./prompts.md)
and [variables](./variables.md) fields are merged with the
root-level fields. When conflicting, the values from
the chosen style will take priority and overwrite
the root-level fields when applicable.

## waku config

=== "Yaml"

    ```yaml
    styles:
      # The key of this map represents
      # the name of the style.
      my-style:
        # Source is the path to the
        # directory containing the files.
        #
        # This has to be relative to the `template.json`
        # file.
        #
        # (required)
        source: ""

        # These fields are optional
        ignore:
        includes:
        labels:
        prompts:
        variables:
    ```

=== "Json"

    ```json
    {
      "styles": {
        // The key of this map represents
        // the name of the style.
        "My Style": {
          // Source is the path to the
          // directory containing the files.
          //
          // This has to be relative to the `template.json`
          // file.
          //
          // (required)
          "source": "",

          // These fields are optional
          "ignore": [],
          "includes": [],
          "labels": [],
          "prompts": [],
          "variables": [],
        }
      }
    }
    ```
