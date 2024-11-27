package git

import (
	"os"
	"os/exec"
)

func Commit(msg string) error {
	cmd := exec.Command("git", "commit", "-m", msg)

	// Redirect stdout and stderr to the console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
