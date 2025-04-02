# Install

There are multiple ways to install Waku on your machine.

## AUR

=== "AUR Helper"

    ```sh
    yay -S waku-bin
    ```

=== "Manually"

    ```sh
    git clone https://aur.archlinux.org/waku-bin.git
    makepkg -si
    ```

## Scoop

```sh
scoop bucket add caffeine-addictt https://github.com/caffeine-addictt/scoop-bucket.git
scoop install waku
```

## Homebrew Tap

```sh
brew install caffeine-addictt/tap/waku
```

## Snapcraft

```sh
sudo snap install waku
```

## Docker

Registries:

- [`caffeinec/waku`](https://hub.docker.com/r/caffeinec/waku)
- [`ghcr.io/caffeine-addictt/waku`](https://github.com/caffeine-addictt/waku/pkgs/container/waku)

**Example usage:**

```sh
docker run -v "$pwd:/app" caffeinec/waku new
```

## Linux packages

Download the `.deb`, `.rpm` or `.apk` packages from the
[release page][releases] and install them with the
appropriate package manager.

After downloading, run:

```sh
dpkg -i waku*.deb
rpm -ivh waku*.rpm
apk add --allow-untrusted waku*.apk
```

## Go

=== "Go install"

    ```sh
    go install github.com/caffeine-addictt/waku@latest
    ```

=== "Manually"

    ```sh
    git clone https://github.com/caffeine-addictt/waku.git
    cd waku
    go install -ldflags="-s -w" .
    ```

Requires Go `1.23+`.

## Verifying installs

### Binaries

All artifacts are checksummed, and the checksum file is signed with [cosign][].

1.  Download the files you want, and the `checksums.txt`, `checksums.txt.pem`
    and `checksums.txt.sig` files from the [releases][] page.

        ```sh
        curl -O 'https://github.com/caffeine-addictt/waku/releases/download/v0.8.1/checksums.txt'
        ```

1.  Verify checksums signature:

    ```bash
    cosign verify-blob \
      --certificate-identity 'https://github.com/caffeine-addictt/waku/.github/workflows/release.yml@refs/tags/v0.8.1' \
      --certificate-oidc-issuer 'https://token.actions.githubusercontent.com' \
      --cert 'https://github.com/caffeine-addictt/waku/releases/download/v0.8.1/checksums.txt.pem' \
      --signature 'https://github.com/caffeine-addictt/waku/releases/download/v0.8.1/checksums.txt.sig' \
      ./checksums.txt
    ```

1.  Verify the SHA256 checksums:

    ```bash
    sha256sum --ignore-missing -c checksums.txt
    ```

### Docker images

Our docker images are signed with [cosign][].

Verify the signature:

```sh
cosign verify \
  --certificate-identity 'https://github.com/caffeine-addictt/waku/.github/workflows/release.yml@refs/tags/v0.8.1' \
  --certificate-oidc-issuer 'https://token.actions.githubusercontent.com' \
  caffeinec/waku
```

[cosign]: https://github.com/sigstore/cosign
[releases]: https://github.com/caffeine-addictt/waku/releases
