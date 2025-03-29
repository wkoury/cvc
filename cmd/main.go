package main

import (
	"cvc/internal/commit"
	"cvc/internal/git"
	"cvc/internal/ollama"
	"cvc/internal/term"
	"flag"
	"fmt"
	"os"
	"slices"
)

func main() {
	flag.Parse()

	args := flag.Args()

	if slices.Contains(args, "ui") {
		commitUI()
	} else {
		commitMessage()
	}
}

// Generate a conventional commit message based on the Git diff using Ollama
func commitMessage() {
	// Get the diff message
	diff, err := git.StagedDiff()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	msg, err := ollama.Request(diff)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(msg)
	os.Exit(0)
}

// Instantiate the UI for making a conventional commit
func commitUI() {
	var commitType commit.ConventionalCommitType
	var breakingChange bool
	var msg string

	term.UI(&commitType, &breakingChange, &msg)

	// We don't need to handle the error here because git commit does that for us
	git.Commit(commit.CommitMsg(commitType, breakingChange, msg))
}
