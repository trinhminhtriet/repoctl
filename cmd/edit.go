package cmd

import (
	"github.com/spf13/cobra"

	"github.com/trinhminhtriet/repoctl/core"
	"github.com/trinhminhtriet/repoctl/core/dao"
)

func editCmd(config *dao.Config, configErr *error) *cobra.Command {
	cmd := cobra.Command{
		Aliases: []string{"e", "ed"},
		Use:     "edit",
		Short:   "Open up repoctl config file",
		Long:    "Open up repoctl config file in $EDITOR.",

		Example: `  # Edit current context
  repoctl edit`,
		Run: func(cmd *cobra.Command, args []string) {
			err := *configErr
			switch e := err.(type) {
			case *core.ConfigNotFound:
				core.CheckIfError(e)
			default:
				runEdit(*config)
			}
		},
		DisableAutoGenTag: true,
	}

	cmd.AddCommand(
		editTask(config, configErr),
		editProject(config, configErr),
	)

	return &cmd
}

func runEdit(config dao.Config) {
	err := config.EditConfig()
	core.CheckIfError(err)
}
