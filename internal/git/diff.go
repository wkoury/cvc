package git

import (
	"bytes"
	"errors"
	"os/exec"
)

func StagedDiff() (string, error) {
	cmd := exec.Command("git", "--no-pager", "diff", "--staged")

	var stdOut bytes.Buffer

	// Redirect stdout to a buffer
	cmd.Stdout = &stdOut

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	ret := stdOut.String()
	if ret == "" {
		return "", errors.New("No staged files")
	}

	return stdOut.String(), nil
}
