package commit

// Given the type, breaking change, and message, return the formatted commit message
func CommitMsg(commitType ConventionalCommitType, isBreakingChange bool, msg string) string {
	ret := string(commitType)

	if isBreakingChange {
		ret += "!"
	}
	ret += ": " + msg

	return ret
}
