# Introduction

Styles are defined through a `template.json` file
usually found in the `root` of the project.

## Validation

You can validate your `template.json` file by running:

```sh
waku check
```

## JSON Schema

Waku also provides a [JSON Schema](https://json-schema.org/)
for better editor support.

```text
https://waku.ngjx.org/static/schema.json
```

Or you can pin a specific version like `v0.5.1`:

```text
https://raw.githubusercontent.com/caffeine-addictt/waku/v0.5.1/www/docs/static/schema.json
```

Simply add the `$schema` property to your `template.json` file:

```json
{
  "$schema": "https://waku.ngjx.org/static/schema.json"
}
```
