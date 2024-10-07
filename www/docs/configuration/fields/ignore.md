# Ignore

Ignore is an array of path globs you want
Waku to ignore when templating.

=== "Yaml"

    ```yaml
    styles:
      my-style:
        ignore:
        # Matching everything.
        - "*"

        # Matching a specific file.
        - "path/to.file"

        # Matching any file starting with `f`
        # in the `path/` directory.
        - "path/f*"

        # Matching all files in the `path/` directory
        # non-recursively.
        - "path/*"

        # Matching all files in the `path/` directory
        # recursively.
        - "path/"
        - "path/**"

        # Placing a `!` infront of any of the above patterns will
        # negate the ignore, explicitly including matches.
        - "!*"
        - "!path/to.file"
        - "!path/f*"
        - "!path/*"
        - "!path/"
        - "!path/**"
    ```

=== "Json"

    ```json
    {
      "ignore": [
        // Matching everything.
        "*",

        // Matching a specific file.
        "path/to.file",

        // Matching any file starting with `f`
        // in the `path/` directory.
        "path/f*",

        // Matching all files in the `path/` directory
        // non-recursively.
        "path/*",

        // Matching all files in the `path/` directory
        // recursively.
        "path/",
        "path/**",

        // Placing a `!` infront of any of the above patterns will
        // negate the ignore, explicitly including matches.
        "!*",
        "!path/to.file",
        "!path/f*",
        "!path/*",
        "!path/",
        "!path/**"
      ]
    }
    ```
