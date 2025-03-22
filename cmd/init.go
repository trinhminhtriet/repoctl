package cmd

import (
	"github.com/spf13/cobra"

	"github.com/trinhminhtriet/repoctl/core"
	"github.com/trinhminhtriet/repoctl/core/dao"
	"github.com/trinhminhtriet/repoctl/core/exec"
)

func initCmd() *cobra.Command {
	var initFlags core.InitFlags

	cmd := cobra.Command{
		Use:   "init",
		Short: "Initialize a repoctl repository",
		Long: `Initialize a repoctl repository.

Creates a new repoctl repository by generating a repoctl.yaml configuration file 
and a .gitignore file in the current directory.`,

		Example: `  # Initialize with default settings
  repoctl init

  # Initialize without auto-discovering projects
  repoctl init --auto-discovery=false

  # Initialize without updating .gitignore
  repoctl init --sync-gitignore=false`,

		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			foundProjects, err := dao.InitMani(args, initFlags)
			core.CheckIfError(err)

			if initFlags.AutoDiscovery {
				exec.PrintProjectInit(foundProjects)
			}
		},
		DisableAutoGenTag: true,
	}

	cmd.Flags().BoolVar(&initFlags.AutoDiscovery, "auto-discovery", true, "automatically discover and add Git repositories to repoctl.yaml")
	cmd.Flags().BoolVarP(&initFlags.SyncGitignore, "sync-gitignore", "g", true, "synchronize .gitignore file")

	return &cmd
}
