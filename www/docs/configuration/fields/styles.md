# Styles

You can specify multiple styles by using the
`styles` field in your `waku config`.

The [setup](./setup.md), [ignore](./ignore.md), [labels](./labels.md)
and [prompts](./prompts.md) fields are merged with the
root-level fields. When conflicting, the values from
the chosen style will take priority and overwrite
the root-level fields when applicable.

## waku config

=== "Yaml"

    ```yaml
    styles:
      my-style:
        styles:
          # The key of this map represents
          # the name of the style.
          My Style:
            # Source is the path to the
            # directory containing the files.
            #
            # This has to be relative to the `template.json`
            # file.
            #
            # (required)
            source: ""

            # These fields are optional
            setup:
            ignore:
            labels:
            prompts:
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
          "setup": {},
          "ignore": [],
          "labels": [],
          "prompts": [],
        }
      }
    }
    ```
