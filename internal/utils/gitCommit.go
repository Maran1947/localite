package utils

import (
	"os"
	"os/exec"
)

func RunGitCommit(message string) error {
	args := []string{"commit","-m",message}
	cmd := exec.Command("git", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}