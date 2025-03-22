---
slug: /
---

# Introduction

`repoctl` is a CLI tool that helps you manage multiple repositories. It's useful when you are working with microservices, multi-project systems, multiple libraries, or just a collection of repositories and want a central place for pulling all repositories and running commands across them.

`repoctl` has many features:

- Declarative configuration
- Clone multiple repositories with a single command
- Run custom or ad-hoc commands across multiple repositories
- Built-in TUI
- Flexible filtering
- Customizable theme
- Auto-completion support
- Portable, no dependencies

## Demo

![demo](/img/demo.gif)

## Example

You specify repositories and commands in a configuration file:

```yaml title="repoctl.yaml"
projects:
  spiko:
    url: https://github.com/trinhminhtriet/spiko.git
    desc: A vim theme editor
    tags: [frontend, node]

  awesome-job-boards:
    url: https://github.com/trinhminhtriet/awesome-job-boards.git
    desc: A simple bash script used to manage boilerplates
    tags: [cli, bash]
    env:
      branch: dev

tasks:
  git-status:
    desc: Show working tree status
    cmd: git status

  git-create:
    desc: Create branch
    spec:
      output: text
    env:
      branch: main
    cmd: git checkout -b $branch
```

Run `repoctl sync` to clone the repositories:

```bash
$ repoctl sync
✓ spiko
✓ rmrfrs

All projects synced
```

Then run commands across all or a subset of the repositories:

```bash
# Target repositories that have the tag 'node'
$ repoctl run git-status --tags node

┌──────────┬─────────────────────────────────────────────────┐
│ Project  │ git-status                                      │
├──────────┼─────────────────────────────────────────────────┤
│ spiko    │ On branch master                                │
│          │ Your branch is up to date with 'origin/master'. │
│          │                                                 │
│          │ nothing to commit, working tree clean           │
└──────────┴─────────────────────────────────────────────────┘

# Target project 'spiko'
$ repoctl run git-create --projects spiko branch=dev

[spiko] TASK [git-create: create branch] *******************

Switched to a new branch 'dev'
```
