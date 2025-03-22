package integration

import (
	"fmt"
	"testing"
)

func TestVersion(t *testing.T) {
	var cases = []TemplateTest{
		{
			TestName:   "Print version when no repoctl config is found",
			InputFiles: []string{},
			TestCmd:    "repoctl --version",
			Ignore:     true,
			WantErr:    false,
		},

		{
			TestName:   "Print version when repoctl config is found",
			InputFiles: []string{"repoctl-advanced/repoctl.yaml"},
			TestCmd:    "repoctl --version",
			Ignore:     true,
			WantErr:    false,
		},
	}

	for i, tt := range cases {
		cases[i].Golden = fmt.Sprintf("version/golden-%d", i)
		cases[i].Index = i
		t.Run(tt.TestName, func(t *testing.T) {
			Run(t, cases[i])
		})
	}
}
