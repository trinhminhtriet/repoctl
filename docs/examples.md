# Examples

This is an example of how to use `repoctl`. Save the following content to a file named `repoctl.yaml` and run `repoctl sync` to clone all repositories. If you already have your own repositories, you can omit the `projects` section.
After setup, you can run any of the [commands](#commands) or check out [git-quick-stats.sh](https://git-quick-stats.sh/) for additional git statistics and run them via `repoctl` for multiple projects.

### Config

```yaml
projects:
  example:
    path: .
    desc: A repoctl example

  spiko:
    path: trinhminhtriet/spiko
    url: https://github.com/trinhminhtriet/spiko.git
    desc: A vim theme editor
    tags: [frontend, node]

  rmrfrs:
    path: trinhminhtriet/rmrfrs
    url: https://github.com/trinhminhtriet/rmrfrs.git
    desc: A highly customizable drag-and-drop grid
    tags: [frontend, lib, node]

  awesome-job-boards:
    url: https://github.com/trinhminhtriet/awesome-job-boards.git
    desc: A simple bash script used to manage boilerplates
    tags: [cli, bash]
    env:
      branch: dev

specs:
  custom:
    output: table
    parallel: true

targets:
  all:
    all: true

themes:
  custom:
    table:
      border:
        around: true
        columns: true
        header: true
        rows: true

tasks:
  git-status:
    desc: show working tree status
    spec: custom
    target: all
    cmd: git status -s

  git-last-commit-msg:
    desc: show last commit
    cmd: git log -1 --pretty=%B

  git-last-commit-date:
    desc: show last commit date
    cmd: |
      git log -1 --format="%cd (%cr)" -n 1 --date=format:"%d  %b %y" \
      | sed 's/ //'

  git-branch:
    desc: show current git branch
    cmd: git rev-parse --abbrev-ref HEAD

  npm-install:
    desc: run npm install in node repos
    target:
      tags: [node]
    cmd: npm install

  git-overview:
    desc: show branch, local and remote diffs, last commit and date
    spec: custom
    target: all
    theme: custom
    commands:
      - task: git-branch
      - task: git-last-commit-msg
      - task: git-last-commit-date
```

## Commands

### List all Projects as Table or Tree:

```bash
$ repoctl list projects

 Project            | Tag                 | Description
--------------------+---------------------+--------------------------------------------------
 example            |                     | A repoctl example
 spiko              | frontend, node      | A vim theme editor
 rmrfrs           | frontend, lib, node | A highly customizable drag-and-drop grid
 awesome-job-boards | cli, bash           | A simple bash script used to manage boilerplates

$ repoctl list projects --tree
┌─ frontend
│  ├─ spiko
│  └─ rmrfrs
└─ awesome-job-boards
```

### Describe Task

```bash
$ repoctl describe task git-overview

Name: git-overview
Description: show branch, local and remote diffs, last commit and date
Theme: custom
Target:
    All: true
    Cwd: false
    Projects:
    Paths:
    Tags:
Spec:
    Output: table
    Parallel: true
    Forks: 4
    IgnoreError: false
    OmitEmptyRows: false
    OmitEmptyColumns: false
Commands:
     - git-branch: show current git branch
     - git-last-commit-msg: show last commit
     - git-last-commit-date: show last commit date
```

### Run a Task Targeting Projects with Tag `node` and Output Table

```bash
$ repoctl run npm-install --tags node

TASK [npm-install: run npm install in node repos] *********************************

spiko |
spiko | up to date, audited 911 packages in 928ms
spiko |
spiko | 71 packages are looking for funding
spiko |   run `npm fund` for details
spiko |
spiko | 15 vulnerabilities (9 moderate, 6 high)
spiko |
spiko | To address issues that do not require attention, run:
spiko |   npm audit fix
spiko |
spiko | To address all issues (including breaking changes), run:
spiko |   npm audit fix --force
spiko |
spiko | Run `npm audit` for details.

TASK [npm-install: run npm install in node repos] *********************************

rmrfrs |
rmrfrs | up to date, audited 960 packages in 1s
rmrfrs |
rmrfrs | 87 packages are looking for funding
rmrfrs |   run `npm fund` for details
rmrfrs |
rmrfrs | 14 vulnerabilities (2 low, 2 moderate, 10 high)
rmrfrs |
rmrfrs | To address all issues possible (including breaking changes), run:
rmrfrs |   npm audit fix --force
rmrfrs |
rmrfrs | Some issues need review, and may require choosing
rmrfrs | a different dependency.
rmrfrs |
rmrfrs | Run `npm audit` for details.
```

### Run Custom Command for All Projects

```bash
$ repoctl exec --all --output table --parallel 'find . -type f | wc -l'

 Project            | Output
--------------------+--------
 example            | 31016
 spiko              | 14444
 rmrfrs           | 16527
 awesome-job-boards | 42
```
