# Quick Start

Here is a quick overview of how to get started with Waku templates.

!!! NOTE

    We will be creating a `waku new template` command
    to generate a new template containing styles soon.

## Create a new template

For demonstration purposes,
let's create a new template containing 1 style called `My Style`.

1.  Create the style directory

    Create a new subdirectory called `my-style`.
    The files you create in this directory
    will be copied and formatted over when `waku new` is ran.

1.  Create a waku `config` file

    === "Yaml"

        Create a `waku.yml` file in the root or subdirectory
        of your project containing the following:

        !!! WARNING

            If you use a subdirectory,
            do not forget to pass the `-d|--directory <path>`
            option when using `waku new`.

        ```yaml
        # yaml-language-server: $schema=https://waku.ngjx.org/static/schema.json

        styles:
          My Style:
            source: style-a
            prompts:
              - key: Description
                ask: A brief description of your project
        ```

    === "Json"

        Create a `waku.json` file in the root or subdirectory
        of your project containing the following:

        !!! WARNING

            If you use a subdirectory,
            do not forget to pass the `-d|--directory <path>`
            option when using `waku new`.

        ```json
        {
          "$schema": "https://waku.ngjx.org/static/schema.json",
          "styles": {
            "My Style": {
              "source": "style-a",
              "prompts": [
                {
                  "key": "Description",
                  "ask": "A brief description of your project"
                }
              ]
            }
          }
        }
        ```
