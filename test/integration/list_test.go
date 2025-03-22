package integration

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	var cases = []TemplateTest{
		// Projects
		{
			TestName:   "List 0 projects",
			InputFiles: []string{"repoctl-empty/repoctl.yaml"},
			TestCmd:    "repoctl list projects",
			WantErr:    false,
		},
		{
			TestName:   "List 0 projects on non-existent tag",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl list projects --tags lala",
			WantErr:    true,
		},
		{
			TestName:   "List 0 projects on 2 non-matching tags",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl list projects --tags frontend,cli",
			WantErr:    false,
		},
		{
			TestName:   "List multiple projects",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl list projects",
			WantErr:    false,
		},
		{
			TestName:   "List only project names and no description/tags",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl list projects --output table --headers project",
			WantErr:    false,
		},
		{
			TestName:   "List projects matching 1 tag",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl list projects --tags frontend",
			WantErr:    false,
		},
		{
			TestName:   "List projects matching multiple tags",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl list projects --tags misc,frontend",
			WantErr:    false,
		},
		{
			TestName:   "List two projects",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl list projects spiko rmrfrs",
			WantErr:    false,
		},
		{
			TestName:   "List projects matching 1 dir",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl list projects --paths frontend",
			WantErr:    false,
		},
		{
			TestName:   "List 0 projects with no matching paths",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl list projects --paths hello",
			WantErr:    true,
		},

		{
			TestName:   "List empty projects tree",
			InputFiles: []string{"repoctl-empty/repoctl.yaml"},
			TestCmd:    "repoctl list projects --tree",
			WantErr:    false,
		},
		{
			TestName:   "List full tree",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl list projects --tree",
			WantErr:    false,
		},
		{
			TestName:   "List tree filtered on tag",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl list projects --tree --tags frontend",
			WantErr:    false,
		},

		// Tags
		{
			TestName:   "List all tags",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl list tags",
			Golden:     "list/tags",
			WantErr:    false,
		},
		{
			TestName:   "List two tags",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl list tags frontend misc",
			Golden:     "list/tags-2-args",
			WantErr:    false,
		},

		// Tasks
		{
			TestName:   "List 0 tasks when no tasks exists ",
			InputFiles: []string{"repoctl-no-tasks/repoctl.yaml"},
			TestCmd:    "repoctl list tasks",
			Golden:     "list/tasks-empty",
			WantErr:    false,
		},
		{
			TestName:   "List all tasks",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl list tasks",
			Golden:     "list/tasks",
			WantErr:    false,
		},
		{
			TestName:   "List two args",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl list tasks fetch status",
			Golden:     "list/tasks-2-args",
			WantErr:    false,
		},
	}

	for i, tt := range cases {
		cases[i].Golden = fmt.Sprintf("list/golden-%d", i)
		cases[i].Index = i

		t.Run(tt.TestName, func(t *testing.T) {
			Run(t, cases[i])
		})
	}
}
