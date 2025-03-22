package integration

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	var cases = []TemplateTest{
		{
			TestName:   "Initialize repoctl in empty directory",
			InputFiles: []string{},
			TestCmd:    "repoctl init --color=false",
			WantErr:    false,
		},

		{
			TestName:   "Initialize repoctl with auto-discovery",
			InputFiles: []string{},
			TestCmd: `
			(mkdir -p rmrfrs && touch rmrfrs/empty);
			(mkdir -p blast && touch blast/empty && cd blast && git init -b main && git remote add origin https://github.com/trinhminhtriet/blast);
			(mkdir -p nested/awesome-job-boards && touch nested/awesome-job-boards/empty && cd nested/awesome-job-boards && git init -b main && git remote add origin https://github.com/trinhminhtriet/awesome-job-boards);
			(mkdir nameless && touch nameless/empty);
			(git init -b main && git remote add origin https://github.com/trinhminhtriet/spiko)
			repoctl init --color=false
			`,
			WantErr: false,
		},

		{
			TestName:   "Throw error when initialize in existing repoctl directory",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl init --color=false",
			WantErr:    true,
		},
	}

	for i, tt := range cases {
		cases[i].Golden = fmt.Sprintf("init/golden-%d", i)
		cases[i].Index = i
		t.Run(tt.TestName, func(t *testing.T) {
			Run(t, cases[i])
		})
	}
}
