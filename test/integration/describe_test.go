package integration

import (
	"fmt"
	"testing"
)

func TestDescribe(t *testing.T) {
	var cases = []TemplateTest{
		// Projects
		{
			TestName:   "Describe 0 projects when there's 0 projects",
			InputFiles: []string{"repoctl-empty/repoctl.yaml"},
			TestCmd:    "repoctl describe projects",
			WantErr:    false,
		},
		{
			TestName:   "Describe 0 projects on non-existent tag",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl describe projects --tags lala",
			WantErr:    true,
		},
		{
			TestName:   "Describe 0 projects on 2 non-matching tags",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl describe projects --tags frontend,cli",
			WantErr:    false,
		},
		{
			TestName:   "Describe all projects",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl describe projects",
			WantErr:    false,
		},
		{
			TestName:   "Describe projects matching 1 tag",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl describe projects --tags frontend",
			WantErr:    false,
		},
		{
			TestName:   "Describe projects matching multiple tags",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl describe projects --tags misc,frontend",
			WantErr:    false,
		},
		{
			TestName:   "Describe 1 project",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl describe projects spiko",
			WantErr:    false,
		},

		// Tasks
		{
			TestName:   "Describe 0 tasks when no tasks exists ",
			InputFiles: []string{"repoctl-no-tasks/repoctl.yaml"},
			TestCmd:    "repoctl describe tasks",
			WantErr:    false,
		},
		{
			TestName:   "Describe all tasks",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl describe tasks",
			WantErr:    false,
		},
		{
			TestName:   "Describe 1 tasks",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl describe tasks status",
			WantErr:    false,
		},
	}

	for i, tt := range cases {
		cases[i].Golden = fmt.Sprintf("describe/golden-%d", i)
		cases[i].Index = i
		t.Run(tt.TestName, func(t *testing.T) {
			Run(t, cases[i])
		})
	}
}
