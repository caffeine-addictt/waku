# Installation

Here you will find all the ways you can install `Waku` on your system.

<!-- prettier-ignore-start -->
<!--toc:start-->

- [Installation](#installation)
  - [AUR](#aur)
  - [Golang](#golang)
  - [Download](#download)
  - [Homebrew](#homebrew)
  - [Choco](#choco)
  - [Docker](#docker)
    - [Our images are hosted on Docker Hub and GitHub](#our-images-are-hosted-on-docker-hub-and-github)
    - [Scripts](#scripts)
      - [Linux/MacOS](#linuxmacos)
      - [Windows](#windows)

<!--toc:end-->
<!-- prettier-ignore-end -->

## AUR

```sh
# Or with your preferred AUR helper
yay -S waku
```

## Golang

```sh
go install github.com/caffeine-addictt/waku@latest
```

## Download

You can also download `Waku` from our
[GitHub release page](https://github.com/caffeine-addictt/waku/releases/latest).

## Homebrew

You can find the homebrew formula
[here](https://github.com/caffeine-addictt/homebrew-tap).

```sh
brew tap caffeine-addictt/tap
brew install caffeine-addictt/tap/waku
```

## Choco

```ps1
choco install waku
```

## Docker

You will need Docker to run `Waku` this way.
You can verify you have Docker installed using the command `docker --version`,
otherwise, you can install Docker from their [docs](https://docs.docker.com/get-started/get-docker/).

### Our images are hosted on Docker Hub and GitHub

> [!NOTE]
> At present, out images do not yet support using
> your system's Git, and thus will only be able to
> access publicly available repositories.
>
> Want to help us out?
> Check out our [Contributing Guide](https://github.com/caffeine-addictt/waku/blob/main/CONTRIBUTING.md).

You will need to mount your current directory to `/app` in the Docker container
for `Waku` to run correctly.

```sh
docker run -v $PWD:/app caffeine/waku:latest
docker run -v $PWD:/app ghcr.io/caffeine-addictt/waku:latest
```

### Scripts

We also provide a set of one-liners to run the `Waku` CLI from docker.

#### Linux/MacOS

```sh
curl -sSL https://github.com/caffeine-addictt/waku/releases/latest/download/waku.sh | bash

# If you prefer to use Waku from GitHub
curl -sSL https://github.com/caffeine-addictt/waku/releases/latest/download/waku.sh | bash -s ghcr
```

#### Windows

```ps1
iwr -useb https://github.com/caffeine-addictt/waku/releases/latest/download/waku.ps1 | iex

# If you prefer to use Waku from GitHub
iwr -useb https://github.com/caffeine-addictt/waku/releases/latest/download/waku.ps1 | iex; Run-Waku "ghcr"
```
