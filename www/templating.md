# Templating

Templating is done through the `template.json` file.
You can have it in your root or sub directory of your
repository.

Do note that having it in a sub directory
will require the `-d|--directory` flag to be set.

<!-- prettier-ignore-start -->
<!--toc:start-->

- [Templating](#templating)
  - [Quick start](#quick-start)
    - [Creating a template](#creating-a-template)
    - [Checking a template](#checking-a-template)
  - [Structure](#structure)
    - [Simple setup](#simple-setup)
    - [Multi style setup](#multi-style-setup)
  - [Fields](#fields)
    - [Setup](#setup)
    - [Ignore](#ignore)
    - [Labels](#labels)
    - [Prompts](#prompts)
    - [Styles](#styles)
      - [Values](#values)

<!--toc:end-->
<!-- prettier-ignore-end -->

## Quick start

### Creating a template

> [!NOTE]
> This is not supported yet.

You can generate a style template by running `waku new template`.

### Checking a template

You can check if a template is valid by running `waku check`.

## Structure

The `template.json` can be structured in 2 ways;
i.e. `name` and `styles` is mutually exclusive.

### Simple setup

Here, `name` is set instead of `styles`.

```json
{
  "$schema": "https://raw.githubusercontent.com/caffeine-addictt/waku/main/schema.json",
  "name": "template name",
  "setup": {
    "linux": "path/to/file.sh",
    "windows": "path/to/file.bat"
  },
  "ignore": [
    "path/to/file",
    "path/to/dir",
    "path/to/files/*",
    "!path/to/files/not/to/ignore"
  ],
  "labels": {
    "my feature": "my color"
  },
  "prompts": {
    "my templated": "Text to prompt user with"
  }
}
```

### Multi style setup

The properties from `"root"` will also be applied
to each style, with each style taking higher priority.

```json
{
  "$schema": "https://raw.githubusercontent.com/caffeine-addictt/waku/main/schema.json",
  "setup": {
    "linux": "path/to/file.sh",
    "windows": "path/to/file.bat"
  },
  "ignore": [
    "path/to/file",
    "path/to/dir",
    "path/to/files/*",
    "!path/to/files/not/to/ignore"
  ],
  "labels": {
    "my feature": "my color"
  },
  "prompts": {
    "my templated": "Text to prompt user with"
  },
  "styles": {
    "my style name": {
      "source": "path/to/style/dir",
      "setup": {
        "linux": "path/to/file.sh",
        "windows": "path/to/file.bat"
      },
      "ignore": [
        "path/to/file",
        "path/to/dir",
        "path/to/files/*",
        "!path/to/files/not/to/ignore"
      ],
      "labels": {
        "my feature": "my color"
      },
      "prompts": {
        "my templated": "Text to prompt user with"
      }
    }
  }
}
```

## Fields

### Setup

> [!NOTE]
> Setup is not supported yet.

`Setup` is a `map` of `name` to `path`.
These executables will be optionally ran
after the template has been generated.

- `Setup.name` can be `linux`, `windows`, `darwin` or `*`.
  `*` is used when the user Os is unknown.
- `Setup.path` is the template-relative path to the executable.

### Ignore

`Ignore` is an array of path globs you do not want templated.
We only support the `*` and `!` special characters.

- `*` is glob for everything
- `path/to/file` is a single file
- `path/to/f*` is a glob for a file
- `path/to/dir/*` is a single dir level glob
- `path/to/dir/**` == `path/to/dir/` is a recursive dir level glob
- `!path/to/file` is a negated ignore

### Labels

> [!NOTE]
> Labels are not supported yet.

Labels are a `map` of `name` to `color`.
Color has to be a hex color code.

```text
Regex
^#(?:[0-9a-fA-F]{3}){1,2}$
```

### Prompts

Prompts are a `map` of `name` to `prompt`.

Name represent the templated value for you files.
I.e. `code` with user responding with `Go` will cause a
templated file `{{CODE}} is cool` -> `Go is cool`.

Prompts are what users are asked with.

### Styles

Styles is a `map` of `name` to `style`.

```json
{
  ...,
  "styles": {
    "My favorite": {
      "source": "style-a",
      ...
    },
    "My friend's": {
      "source": "style-b",
      ...
    },
    ...
  }
}
```

- **`my-templates/`**
  - **`template.json`**
  - **`style-a/`**
    - `file1`
    - `file2`
    - `...`
  - **`style-b/`**
    - `file1`
    - `file2`
    - `...`
  - `...`

#### Values

- `style.source` is the path to the style directory to use.
- [`style.setup`](#setup)
- [`style.ignore`](#ignore)
- [`style.labels`](#labels)
- [`style.prompts`](#prompts)
