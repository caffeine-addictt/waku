# Introduction

Styles are defined through a `configuration` file
usually found in the `root` of the project.

## Validation

You can validate your `waku.yml` file by running:

```sh
waku check
```

## JSON Schema

Waku also provides a [JSON Schema](https://json-schema.org/)
for better editor support.

```text
https://waku.ngjx.org/static/schema.json
```

Or you can pin a specific version like `v0.8.1`:

```text
https://raw.githubusercontent.com/caffeine-addictt/waku/v0.8.1/www/docs/static/schema.json
```

Simply add the `$schema` property to your `configuration` file:

=== "Yaml"

    Add the following to your `waku.yml` file as a comment

    ```yaml
    # yaml-language-server: $schema=https://waku.ngjx.org/static/schema.json
    ```

=== "Json"

    ```json
    {
      "$schema": "https://waku.ngjx.org/static/schema.json"
    }
    ```

## Filenames

Here are the default filenames we look for:

- `waku.yml`
- `waku.yaml`
- `waku.json`
- `template.yml`
- `template.yaml`
- `template.json`
- `.waku.yml`
- `.waku.yaml`
- `.waku.json`
- `.template.yml`
- `.template.yaml`
- `.template.json`
