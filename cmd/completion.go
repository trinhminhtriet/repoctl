package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/trinhminhtriet/repoctl/core"
)

func completionCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "completion <bash|zsh|fish|powershell>",
		Short: "Generate completion script",
		Long: `To load completions:
Bash:

  $ source <(repoctl completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ repoctl completion bash > /etc/bash_completion.d/repoctl
  # macOS:
  $ repoctl completion bash > /usr/local/etc/bash_completion.d/repoctl

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ repoctl completion zsh > "${fpath[1]}/_repoctl"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ repoctl completion fish | source

  # To load completions for each session, execute once:
  $ repoctl completion fish > ~/.config/fish/completions/repoctl.fish

PowerShell:

  PS> repoctl completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> repoctl completion powershell > repoctl.ps1
  # and source this file from your PowerShell profile.
		`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run:                   generateCompletion,
		DisableAutoGenTag:     true,
	}

	return &cmd
}

func generateCompletion(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "bash":
		err := cmd.Root().GenBashCompletion(os.Stdout)
		core.CheckIfError(err)
	case "zsh":
		err := cmd.Root().GenZshCompletion(os.Stdout)
		core.CheckIfError(err)
	case "fish":
		err := cmd.Root().GenFishCompletion(os.Stdout, true)
		core.CheckIfError(err)
	case "powershell":
		err := cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		core.CheckIfError(err)
	}
}
