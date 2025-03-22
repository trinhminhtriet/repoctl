package integration

import (
	"fmt"
	"testing"
)

func TestExec(t *testing.T) {
	var cases = []TemplateTest{
		{
			TestName:   "Should fail to exec when no configuration file found",
			InputFiles: []string{},
			TestCmd: `
				repoctl exec --all -o table ls
			`,
			WantErr: true,
		},

		{
			TestName:   "Should exec in zero projects",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
				repoctl sync
				repoctl exec -o table ls
			`,
			WantErr: true,
		},

		{
			TestName:   "Should exec in all projects",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
				repoctl sync
				repoctl exec --all -o table ls
			`,
			WantErr: false,
		},

		{
			TestName:   "Should exec when filtered on project name",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
				repoctl sync
				repoctl exec -o table --projects spiko ls
			`,
			WantErr: false,
		},

		{
			TestName:   "Should exec when filtered on tags",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
				repoctl sync
				repoctl exec -o table --tags frontend ls
			`,
			WantErr: false,
		},

		{
			TestName:   "Should exec when filtered on cwd",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
				repoctl sync
				cd awesome-job-boards
				repoctl exec -o table --cwd pwd
			`,
			WantErr: false,
		},

		{
			TestName:   "Should dry run exec",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
				repoctl sync
				repoctl exec -o table --dry-run --projects awesome-job-boards pwd
			`,
			WantErr: false,
		},
	}

	for i, tt := range cases {
		cases[i].Golden = fmt.Sprintf("exec/golden-%d", i)
		cases[i].Index = i
		t.Run(tt.TestName, func(t *testing.T) {
			Run(t, cases[i])
		})
	}
}
