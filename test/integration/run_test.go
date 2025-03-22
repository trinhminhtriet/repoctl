package integration

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	var cases = []TemplateTest{
		{
			TestName:   "Should fail to run when no configuration file found",
			InputFiles: []string{},
			TestCmd: `
			repoctl run pwd --all
			`,
			WantErr: true,
		},

		{
			TestName:   "Should run in zero projects",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
			repoctl sync
			repoctl run pwd -o table
			`,
			WantErr: true,
		},

		{
			TestName:   "Should run in all projects",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
			repoctl sync
			repoctl run --all pwd
			`,
			WantErr: false,
		},

		{
			TestName:   "Should run when filtered on project",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
			repoctl sync
			repoctl run -o table --projects spiko pwd
			`,
			WantErr: false,
		},

		{
			TestName:   "Should run when filtered on tags",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
			repoctl sync
			repoctl run -o table --tags frontend pwd
			`,
			WantErr: false,
		},

		{
			TestName:   "Should run when filtered on cwd",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
			repoctl sync
			cd awesome-job-boards
			repoctl run -o table --cwd pwd
			`,
			WantErr: false,
		},

		{
			TestName:   "Should run on default tags",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
			repoctl sync
			repoctl run -o table default-tags
			`,
			WantErr: false,
		},

		{
			TestName:   "Should run on default projects",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
			repoctl sync
			repoctl run -o table default-projects
			`,
			WantErr: false,
		},

		{
			TestName:   "Should print table when output set to table in task",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
			repoctl sync
			repoctl run default-output -p rmrfrs
			`,
			WantErr: false,
		},

		{
			TestName:   "Should dry run",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
			repoctl sync
			repoctl run --dry-run --projects awesome-job-boards -o table pwd
			`,
			WantErr: false,
		},

		{
			TestName:   "Should run multiple commands",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
			repoctl sync
			repoctl run pwd multi -o table --all
			`,
			WantErr: false,
		},

		{
			TestName:   "Should run sub-commands",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml", "repoctl-advanced/.gitignore"},
			TestCmd: `
			repoctl sync
			repoctl run submarine --all
			`,
			WantErr: false,
		},
	}

	for i, tt := range cases {
		cases[i].Golden = fmt.Sprintf("run/golden-%d", i)
		cases[i].Index = i
		t.Run(tt.TestName, func(t *testing.T) {
			Run(t, cases[i])
		})
	}
}
