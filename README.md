# 🚀 repoctl

```text
                           _   _
 _ __ ___ _ __   ___   ___| |_| |
| '__/ _ \ '_ \ / _ \ / __| __| |
| | |  __/ |_) | (_) | (__| |_| |
|_|  \___| .__/ \___/ \___|\__|_|
         |_|
```

🚀 repoctl – A powerful CLI tool to manage multiple Git repositories effortlessly. Sync, pull, and run commands! 🎯

## ✨ Features

- **📜 Declarative Configuration** – Define repositories and actions in a simple configuration file.
- **⚡ Clone Multiple Repositories** – Clone and manage multiple repositories with a single command.
- **🔧 Run Commands Across Repos** – Execute custom or ad-hoc commands across all or selected repositories.
- **🖥️ Built-in TUI** – Intuitive terminal user interface for easy navigation and control.
- **🎯 Flexible Filtering** – Select repositories based on patterns, groups, or custom criteria.
- **🎨 Customizable Theme** – Personalize the UI with themes to match your preferences.
- **⌨️ Auto-Completion Support** – Bash, Zsh, and Fish shell completions for a seamless experience.
- **📦 Portable & No Dependencies** – A single binary with zero external dependencies.

## 🚀 Installation

Download from [latest releases ](https://github.com/trinhminhtriet/repoctl/releases)

### Build from source

```bash
git clone
cd repoctl
make build

./dist/repoctl --version
./dist/repoctl --help
cp ./dist/repoctl /usr/local/bin/repoctl
```

## 💡 Usage

```bash
# List all projects
repoctl list projects

# Count number of files in each project in parallel
repoctl exec --all --output table --parallel 'find . -type f | wc -l'

# Start TUI
repoctl tui
```

## 🤝 How to contribute

We welcome contributions!

- Fork this repository;
- Create a branch with your feature: `git checkout -b my-feature`;
- Commit your changes: `git commit -m "feat: my new feature"`;
- Push to your branch: `git push origin my-feature`.

Once your pull request has been merged, you can delete your branch.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
