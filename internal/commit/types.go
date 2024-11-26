package commit

// Conventional commit types
type ConventionalCommitType string

const (
	ConventionalCommitTypeFeat     = ConventionalCommitType("feat")
	ConventionalCommitTypeFix      = ConventionalCommitType("fix")
	ConventionalCommitTypeBuild    = ConventionalCommitType("build")
	ConventionalCommitTypeChore    = ConventionalCommitType("chore")
	ConventionalCommitTypeCI       = ConventionalCommitType("ci")
	ConventionalCommitTypeDocs     = ConventionalCommitType("docs")
	ConventionalCommitTypeStyle    = ConventionalCommitType("style")
	ConventionalCommitTypeRefactor = ConventionalCommitType("refactor")
	ConventionalCommitTypePerf     = ConventionalCommitType("perf")
	ConventionalCommitTypeTest     = ConventionalCommitType("test")
)

// An ordered slice of all the conventional commits we want in our picker
var (
	AllConventionalCommitTypes = []string{
		string(ConventionalCommitTypeFix),
		string(ConventionalCommitTypeFeat),
		string(ConventionalCommitTypeRefactor),
		string(ConventionalCommitTypeTest),
		string(ConventionalCommitTypeChore),
		string(ConventionalCommitTypeBuild),
		string(ConventionalCommitTypeCI),
		string(ConventionalCommitTypeDocs),
		string(ConventionalCommitTypeStyle),
		string(ConventionalCommitTypePerf),
	}
)
