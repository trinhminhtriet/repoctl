# Installation

`repoctl` is available on Linux and Mac, with partial support for Windows.

* Binaries are available on the [release](https://github.com/trinhminhtriet/repoctl/releases) page

* via cURL (Linux & macOS)
  ```bash
  curl -sfL https://raw.githubusercontent.com/trinhminhtriet/repoctl/main/install.sh | sh
  ```

* via Homebrew
  ```bash
  brew tap trinhminhtriet/repoctl
  brew install repoctl
  ```

* via Arch
  ```sh
  pacman -S repoctl
  ```

* via Nix
  ```sh
  nix-env -iA nixos.repoctl
  ```

* via Go
  ```bash
  go get -u github.com/trinhminhtriet/repoctl
  ```

## Building From Source

1. Clone the repo
2. Build and run the executable

    ```bash
    make build && ./dist/repoctl
    ```
