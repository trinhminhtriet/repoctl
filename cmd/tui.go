package cmd

import (
	"github.com/spf13/cobra"
	"github.com/trinhminhtriet/repoctl/core"
	"github.com/trinhminhtriet/repoctl/core/dao"
	"github.com/trinhminhtriet/repoctl/core/tui"
)

func tuiCmd(config *dao.Config, configErr *error) *cobra.Command {
	var tuiFlags core.TUIFlags

	cmd := cobra.Command{
		Use:     "tui",
		Aliases: []string{"gui"},
		Short:   "TUI",
		Long:    `Run TUI`,
		Example: `  # Open tui
  repoctl tui`,
		Run: func(cmd *cobra.Command, args []string) {
			core.CheckIfError(*configErr)
			reloadChanged := cmd.Flags().Changed("reload-on-change")
			reload := config.ReloadTUI
			if reloadChanged {
				reload = &tuiFlags.Reload
			}

			tui.RunTui(config, tuiFlags.Theme, *reload)
		},
		DisableAutoGenTag: true,
	}

	cmd.PersistentFlags().StringVar(&tuiFlags.Theme, "theme", "default", "set theme")
	err := cmd.RegisterFlagCompletionFunc("theme", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if *configErr != nil {
			return []string{}, cobra.ShellCompDirectiveDefault
		}

		names := config.GetThemeNames()

		return names, cobra.ShellCompDirectiveDefault
	})
	core.CheckIfError(err)

	cmd.Flags().BoolVarP(&tuiFlags.Reload, "reload-on-change", "r", false, "reload repoctl on config change")

	return &cmd
}
