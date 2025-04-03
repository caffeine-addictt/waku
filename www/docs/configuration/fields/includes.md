# Includes

The `includes` field in the `waku config` allows you to specify additional directories
to include in your templating process. This can be helpful when you want to include
extra files or directories that should be copied over, in addition to the default
templates.

Each `include` can optionally specify which files should be ignored and where
the included files should be placed in the target structure.

## waku config

The `includes` field supports 2 formats:

- **An array of strings**
- **An array of objects**

### Paths Precedence: Style > Includes

When the resolved resource path from the `style` and `include` collide,
the `style` path takes precedence.

For example, when `shared/README.md` and `my-style/README.md` both exist,
`my-style/README.md` will take precedence every time.

!!! INFO

    Paths specified within the `includes` field, except the `dir` field,
    is **relative to the parent directory of your `waku config` file**
    and not the style source directory.

    Include sources also cannot self-reference the style source or any of its
    subdirectories.

### Format 1: Array of Strings

In this format, you can provide a simple list of
directory paths to include.

=== "Yaml"

    ```yaml
    styles:
      my-style:
        includes:
          - assets
          - shared-files
    ```

=== "Json"

    ```json
    {
      "styles": {
        "My Style": {
          "includes": [
            "assets",
            "shared-files"
          ]
        }
      }
    }
    ```

### Format 2: Array of Objects

In this format, you can specify more advanced options for each include,
such as `source`, `ignore`, and `dir`.

=== "Yaml"

    ```yaml
    styles:
      my-style:
        includes:
          # Source is the path to the
          # directory containing the files.
          #
          # This has to be relative to the parent directory of
          # your waku config file, and cannot self-reference
          # the style source or any of its subdirectories.
          #
          # (required)
          - source: ""

          # Dir is the directory to place the included files in
          # when templating.
          #
          # This is relative to the parent directory of your
          # new templated project.
          #
          # (optional)
            dir: ""

          # Ignore is a list of paths or patterns to ignore.
          # See the ignore field for more info.
          #
          # (optional)
            ignore: []
    ```

=== "Json"

    ```json
    {
      "styles": {
        "My Style": {
          "includes": [
            {
              // Source is the path to the
              // directory containing the files.
              //
              // This has to be relative to the parent directory of
              // your waku config file, and cannot self-reference
              // the style source or any of its subdirectories.
              //
              // (required)
              "source": "",

              // Dir is the directory to place the included files in
              // when templating.
              //
              // This is relative to the parent directory of your
              // new templated project.
              //
              // (optional)
              "dir": "",

              // Ignore is a list of paths or patterns to ignore.
              // See the ignore field for more info.
              //
              // (optional)
              "ignore": []
            }
          ]
        }
      }
    }
    ```
