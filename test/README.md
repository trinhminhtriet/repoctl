# Test

`repoctl` currently only has integration tests, which require `docker` to run. This is because `repoctl` mainly interacts with the filesystem, and whilst there are ways to mock the filesystem, it's simply easier (and fast enough) to spin up a `docker` container and do the work there.

The tests are based on something called "golden files", which are the expected output of the tests. It serves the benefit of working as documentation as well, since it becomes easy to see the desired output of the different `repoctl` commands.

There's some helpful scripts in the `scripts` directory that can be used to test and debug `repoctl`. These scripts should be run from the project directory.

The Docker test container includes a script `git` which only creates the project directories, it doesn't clone the actual repositories.

## Directory Structure

```sh
.
├── fixtures    # files needed for testing purposes
├── images      # docker images used for testing and development
├── integration # integration tests and golden files
├── scripts     # scripts for development and testing
└── tmp         # docker mounted volume that you can preview test output
```

## Prerequisites

- [docker](https://docs.docker.com/get-docker/)
- [golangci-lint](https://golangci-lint.run/usage/install/)

## Testing & Development

Checkout the below commands and the [Makefile](../Makefile) to test/debug `repoctl`.

```sh
# Run tests
./test/scripts/test

# Run specific tests, print stdout and build repoctl
./test/scripts/test --debug --build --run TestInitCmd

# Update Golden Files
./test/scripts/test -u

# Start an interactive shell inside docker
./test/scripts/exec --shell bash|zsh|fish

# Debug completion
repoctl __complete list tags --projects ""

# Stand in _example directory
(cd ../ && make build-and-link && cd - && repoctl run status --cwd)
```
