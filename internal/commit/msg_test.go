package commit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommitMsg(t *testing.T) {
	testCases := []struct {
		name             string
		commitType       ConventionalCommitType
		isBreakingChange bool
		msg              string
		expected         string
	}{
		{
			"regular fix",
			ConventionalCommitTypeFix,
			false,
			"resolve bug",
			"fix: resolve bug",
		},
		{
			"regular feat",
			ConventionalCommitTypeFeat,
			false,
			"add function",
			"feat: add function",
		},
		{
			"breaking fix",
			ConventionalCommitTypeFix,
			true,
			"resolve bug",
			"fix!: resolve bug",
		},
		{
			"breaking feat",
			ConventionalCommitTypeFeat,
			true,
			"add function",
			"feat!: add function",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, CommitMsg(tc.commitType, tc.isBreakingChange, tc.msg))
		})
	}
}
