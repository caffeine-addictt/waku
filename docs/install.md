# Installation

Here you will find all the ways you can install `Waku` on your system.

<!-- prettier-ignore-start -->
<!--toc:start-->

- [Installation](#installation)
  - [AUR](#aur)
  - [Golang](#golang)
  - [Download](#download)
  - [Snapcraft](#snapcraft)
  - [Homebrew](#homebrew)
  - [Scoop](#scoop)
  - [Choco](#choco)
  - [Docker](#docker)
    - [Our images are hosted on Docker Hub and GitHub](#our-images-are-hosted-on-docker-hub-and-github)
    - [Scripts](#scripts)
      - [Linux/MacOS](#linuxmacos)
      - [Windows](#windows)
  - [Verify](#verify)
    - [1. Download the checksums.txt file](#1-download-the-checksumstxt-file)
    - [2. Verify checksums signature](#2-verify-checksums-signature)
    - [3. Verify the SHA256 sums match](#3-verify-the-sha256-sums-match)

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

## Snapcraft

```sh
sudo snap install waku --classic
```

## Homebrew

You can find the homebrew formula
[here](https://github.com/caffeine-addictt/homebrew-tap).

```sh
brew tap caffeine-addictt/tap
brew install caffeine-addictt/tap/waku
```

## Scoop

```sh
scoop bucket add caffeine https://github.com/caffeine-addictt/scoop-bucket.git
scoop install <name>
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
docker run -v $PWD:/app caffeinec/waku:latest
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

## Verify

### 1. Download the checksums.txt file

Either from our [releases page](https://github.com/caffeine-addictt/waku/releases/latest)
or by running the following command:

```sh
curl -sSL https://github.com/caffeine-addictt/waku/releases/latest/download/checksums.txt -o checksums.txt
```

### 2. Verify checksums signature

We sign our checksums with [Cosign](https://github.com/sigstore/cosign)

```sh
cosign verify-blob \
  --certificate-identity 'https://github.com/caffeine-addictt/waku/.github/workflows/release.yml@refs/tags/v0.3.0' \
  --certificate-oidc-issuer 'https://token.actions.githubusercontent.com' \
  --cert 'https://github.com/caffeine-addictt/waku/releases/download/v0.3.0/checksums.txt.pem' \
  --signature 'https://github.com/caffeine-addictt/waku/releases/download/v0.3.0/checksums.txt.sig' \
  ./checksums.txt
```

### 3. Verify the SHA256 sums match

```sh
sha256sum --ignore-missing -c checksums.txt
```
