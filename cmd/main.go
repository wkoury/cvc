package main

import (
	"cvc/internal/commit"
	"cvc/internal/git"
	"cvc/internal/term"
)

func main() {
	var commitType commit.ConventionalCommitType
	var breakingChange bool
	var msg string

	term.UI(&commitType, &breakingChange, &msg)

	// We don't need to handle the error here because `git commit does that for us`
	git.Commit(commit.CommitMsg(commitType, breakingChange, msg))
}
