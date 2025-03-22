# Usage

## Create a New Repoctrl Repository

Run the following command inside a directory containing your `git` repositories:

```bash
$ repoctl init
```

This will generate **two** files:

- `repoctl.yaml`: Contains projects and custom tasks. Any subdirectory that has a `.git` directory will be included (add the flag `--auto-discovery=false` to turn off this feature)
- `.gitignore`: Includes the projects specified in `repoctl.yaml` file. To opt out, use `repoctl init --vcs=none`.

It can be helpful to initialize the `repoctl` repository as a git repository so that anyone can easily download the `repoctl` repository and run `repoctl` sync to clone all repositories and get the same project setup as you.

## Run Some Commands

```bash
# List all projects
$ repoctl list projects

# Count number of files in each project in parallel
$ repoctl exec --all --output table --parallel 'find . -type f | wc -l'
```

Next up:

- [Some more examples](/examples)
- [Familiarize yourself with the repoctl.yaml config](/config)
- [Checkout repoctl commands](/commands)
