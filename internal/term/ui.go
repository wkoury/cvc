// package term contains the command line interface
package term

import (
	"cvc/internal/commit"
	"cvc/internal/git"

	"github.com/rivo/tview"
)

func UI() {
	var commitType commit.ConventionalCommitType
	breakingChange := false
	msg := ""

	commitMsgInputWidth := 80
	commitMsgInputHeight := 3
	commitMsgMaxLength := commitMsgInputWidth * commitMsgInputHeight

	app := tview.NewApplication()
	form := tview.NewForm().
		AddDropDown("Commit type", commit.AllConventionalCommitTypes, 0, func(option string, optionIndex int) {
			commitType = commit.ConventionalCommitType(option)
		}).
		AddCheckbox("Breaking change?", false, func(checked bool) { breakingChange = checked }).
		AddTextArea("Commit message", "", commitMsgInputWidth, commitMsgInputHeight, commitMsgMaxLength, func(text string) { msg = text }).
		AddButton("Commit", func() {
			app.Stop()
			if err := git.Commit(commit.CommitMsg(commitType, breakingChange, msg)); err != nil {
				panic(err)
			}
		})

	form.SetBorder(true).SetTitle("Conventional commit").SetTitleAlign(tview.AlignCenter)
	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
