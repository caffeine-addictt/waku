# **Contributing**

By contributing to Waku, you agree to abide by our
[code of conduct](https://github.com/caffeine-addictt/waku?tab=coc-ov-file#readme).

## Prerequisites

- [Go 1.23+][Go install]
- [GNU/Make][]
- [Prettier][]
- [Python 3.11+][Python] (docs)

## Building

You can build Waku with [GNU/Make][],
simply run the following commands:

### Go

```sh
make build
```

### Docker

```sh
make build/docker
```

### Docs

```sh
make build/docs
```

## Testing

Running the following will run all tests
as well as check for vulnerabilities.

```sh
make test
```

During development, you can also run the following:

### Go

This will run the Waku CLI.

```sh
go run .
```

### Docs

This will start a development server for documentation
on [`http://localhost:8000`](http://localhost:8000).

```sh
make docs
```

## Creating commits

Commit messages should conform to [Conventional commits][].

We also ask that your code is formatted before committing
by running:

```sh
make lint
make fmt
```

## Submitting a Pull Request

Push your changes to your `waku` fork and create a Pull Request
against the main branch. Please include a clear and concise description
of your changes and why they are necessary. Also ensure that your Pull Request
title conforms to [Conventional commits][] and that you have incremented version
numbers according to [SemVer][] by running:

```sh
make bump version=x.x.x
```

## Creating an issue

Please ensure that there is no similar issue already open before
creating a new one.

If not, you can choose a relevant issue template from the [list](https://github.com/caffeine-addictt/waku/issues/new/choose).
Providing as much information as possible will make it easier for us to help
resolve your issue.

## Financial contributions

You can consider sponsoring Waku.
See [this page](https://waku.ngjx.org/sponsors) for more details.

[Go install]: https://go.dev/doc/install
[GNU/Make]: https://www.gnu.org/software/make/#download
[Prettier]: https://prettier.io/
[Python]: https://www.python.org
[Conventional commits]: https://www.conventionalcommits.org
[SemVer]: https://semver.org
