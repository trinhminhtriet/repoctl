# Changelog

## 0.1.0

### Features

- Added a sub-command to launch a TUI
- Added `--forks` flag to limit parallel task execution [#74](https://github.com/trinhminhtriet/repoctl/issues/74)
- Added `--target` specification from flags [#82](https://github.com/trinhminhtriet/repoctl/issues/82)
- Added `--spec` specification from flags
- Added `--ignore-sync-state` flag to `repoctl sync` to ignore `sync` status set projects [#83](https://github.com/trinhminhtriet/repoctl/issues/83)
- Added `--tags-expr` flag for complex tag filtering expressions (e.g., (active || git) targets projects with either active or git tag) [#85](https://github.com/trinhminhtriet/repoctl/issues/85)
- Added `--sync-gitignore` flag to opt out of `.gitignore` file modifications [#87](https://github.com/trinhminhtriet/repoctl/issues/87)
- Added `tty` attribute to tasks which will replace the command and allow attaching to docker containers

### Fixes

- Fixed `repoctl init` behavior when root directory contains `.git` [#78](https://github.com/trinhminhtriet/repoctl/issues/78)
- Fixed `repoctl sync` execution when running `repoctl init` with remotes [#84](https://github.com/trinhminhtriet/repoctl/issues/84)
- Fixed table column truncation when output exceeds terminal width

### Misc

- Changed filtering tags/paths behavior to use intersection instead of union
- Changed default shell from `sh` to `bash`
- Improved multiple task execution by treating them as sub-commands for cleaner output
- Renamed `--no-color` flag to `--color`
- Changed output `text` to `stream` for all outputs (`flags`, `themes`, and `spec`)
- Updated theme configuration system
- Enhanced remote management: `repoctl` now removes git remotes if specified via global field `sync_remotes` config or flag `--sync-remotes`

## 0.25.0

### Features

- Add more box styles to table and tree output

### Misc

- Update golang to 1.20.0

## 0.24.0

### Features

- Add ability to create/sync remotes

## 0.23.0

### Features

- Add option `--ignore-non-existings` to ignore projects that don't exist
- Add flag `--ignore-errors` to ignore errors

## 0.22.0

### Features

- Add filter options to sub-command sync [#52](https://github.com/trinhminhtriet/repoctl/pull/52)
- Add check sub-command to validate repoctl config
- Add option to disable spinner when running tasks [#54](https://github.com/trinhminhtriet/repoctl/pull/54)

### Fixes

- Fix wrongly formatted YAML for init command

## 0.21.0

### Features

- Add path and url env to project clone command

## 0.20.1

### Fixes

- Fix evaluate env for MANI_CONFIG and MANI_USER_CONFIG
- Fix parallel sync, limit to 20 projects at a time

### Changes

- Use `repoctl --version` flag instead of `repoctl version`

## 0.20.0

A lot of refactoring and some new features added. There's also some breaking changes, notably how themes work.

### Features

- Add option to skip sync on projects by setting `sync` property to `false`
- Add flag to disable colors and respect NO_COLOR env variable when set
- Add env variables MANI_CONFIG and MANI_USER_CONFIG that checks main config and user config
- Add desc of tasks when auto-completing
- Add man page generation
- [BREAKING CHANGE]: Major theme overhaul, allow granular theme modification

### Fix

- Don't automatically create the `$XDG_CONFIG_HOME/repoctl/config.yaml` file
- Fix overriding spec data (parallel and omit-empty-columns) with flags
- Fix when initializing repoctl with multiple repos having the same name [#30](https://github.com/trinhminhtriet/repoctl/issues/30), thanks to https://github.com/stessaris for finding the bug
- Omit empty now checks all command outputs, and omits iff all of them are empty
- Start spinner after 500 ms to avoid flickering when running commands which take less than 500 ms to execute

### Changes

- [BREAKING CHANGE]: Remove no-headers flag
- [BREAKING CHANGE]: Remove no-borders flag and enable it to be configurable via theme
- [BREAKING CHANGE]: Removed default env variables that was set previously (MANI_PROJECT_PATH, .etc)
- Remove some acceptable repoctl config filenames (notably, those that do not end in .yaml/.yml)
- Update task and project describe
- Improve error messages

### Internal

- A lot of refactoring
  - Rework exec.Cmd
  - Remove aurora color library dependency and use the one provided by go-pretty

## v0.12.2

### Fixes

- Allow placing repoctl config inside one of directories of a repoctl project when syncing

## v0.12.0

### Features

- Add option to omit empty results
- Add --vcs flag to repoctl init to choose vcs
- Add default import from user config directory
- [BREAKING CHANGE]: Add spec property to allow reusing common properties
- Add target property to allow reusing common properties

### Fixes

- Fix header bug in run print when task has both commands and cmd
- Fix `repoctl edit` to run even if config file is malformed (wrong YAML syntax)

### Misc

- [BREAKING CHANGE]: Move tree feature to list projects as a flag instead of it being a special sub-command
- [BREAKING CHANGE]: Rename flag --all-projects to all
- Remove legacy code related to Dirs entity
- Change default value of --parallel flag to false when syncing
- Allow omitting the -c when specifying shell for bash, zsh, sh, node and python

## v0.11.1

### Fixes

- Use syncmap to allow safe concurrent writes when running `repoctl sync` in parallel, previously there was a race condition that occurred when cloning many repos

### Features

- Add `env` property to projects to enable project specific variables

## v0.10.0

### Features

- Add ability to import projects, tasks and themes
- Possible to run tasks in parallel now per each project
- Add sub-commands project/task to edit command to open editor at line corresponding to project/task
- Add edit flag to describe/run sub-commands to open up editor
- Sync projects in parallel by default and add flag serial to opt out
- Add support for referencing commands in Commands property
- Run commands in serial, if one fails, dont run other tasks
- Add directory entity, similar to project, just without a url/clone property

### Misc

- Add new acceptable filenames Manifile, Manifile.yaml, Manifile.yml
- Don't create .gitignore if no projects with url exists on repoctl init/sync
- List tags now shows associated dirs/projects
- If user uses a cwd/tag/project/dir flag, then disable task targets
- [BREAKING CHANGE:] A lot of syntax changes, use object notation instead of array list for projects, themes and tasks

## v0.6.1

### Features

- Add dirs filtering property to commands struct

### Fixes

- Correct project path in gitignore file when running repoctl init

### Misc

- Update help text for dirs flag

## v0.6.0

### Features

- New tree command that list contents of projects in a tree-like format
- Add filtering on directory for tree/list/describe/run/exec cmd
- Add global environment variables
- Add describe flag to run cmd to suppress command information
- Add sub-commands
- Add possibility to run multiple commands from cli
- Add default tags/projects/output to tasks
- Add new table style that can be configured only from repoctl config
- Add progress spinner for run/exec cmd

### Misc

- [BREAKING CHANGE]: Renamed args in command block to env
- [BREAKING CHANGE]: Renamed commands in root block to tasks
- Environment variables now support shell execution
- Rename flag format to output when listing

## v0.5.1

### Fixes

- Fix auto-complete for flag format in list command

## v0.5.0

### Features

- Add Repoctrl environment variable that is cwd of the current context repoctl.yaml file
- Add repoctl edit command which opens repoctl.yaml in preferred editor
- Add describe cmd, display commands and projects in detail
- Append default shell to commands
- Add output formats table, markdown and html
- Add no-borders, no-headers flags to print
- Allow users to specify headers to be printed in list command
- Sync creates gitignore file if not found
- Use CLI spinner when syncing projects
- Update info cmd to print git version

### Fixes

- Output args at top for run commands instead of for each run
- Output error message when running commands in non-repoctl directory that require repoctl config

### Misc

- Refactor and make code more DRY
- Refactor list and describe cmd to use sub-commands
- With no projects to sync, output helpful message: "No projects to sync"
- With all projects synced, output helpful message: "All projects synced"

## v0.4.0

### Features

- Allow users to set global and command level shell commands

## v0.3.0

### Features

- Add support for running from nested sub-directories
- Add info sub-command that shows which configuration file is being used
- Add flag to point to config file
- Accept different config names (.repoctl, .repoctl.yaml, .repoctl.yml, repoctl.yaml, repoctl.yml)
- Add new command exec to run arbitrary command
- Add config flag
- Add first argument to init should be path, if empty, current dir
- Add completion for all commands bash
- Update auto-discovery to equal true by default
- Add option to filter list command on tags and projects
- Add Nicer output on failed git sync
- Add cwd flag to target current directory
- Add comment section in .gitignore so users can modify the gitignore without repoctl overwriting all parts
- Improved listing for projects/tags

### Fixes

- Fix crashing on not found config file
- Check possible, non-handled nil/err values
- Don't add project to gitignore if doesn't have a url
- Remove path if path is same as name
- Fix gitignore sync, removing old entries
- Fix broken init command
- Fix so path accepts environment variables
- Fix auto-complete when not in repoctl directory

### Misc

- Update golang version and dependencies
- Add integration tests
