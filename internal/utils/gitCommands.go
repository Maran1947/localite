package utils

import (
	"os"
	"os/exec"
	"strings"
)

func RunGitAdd() error {
	args := []string{"add","."}
	cmd := exec.Command("git", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func RunGitDiff(flags string) (string, error) {
	err := RunGitAdd()
	if err != nil {
		return "", err
	}

	cmdArgs := []string{"diff","--cached"}
	if flags != "" {
        cmdArgs = append(cmdArgs, strings.Fields(flags)...)
    }

	cmd := exec.Command("git", cmdArgs...)

	output, err := cmd.Output()
    if err != nil {
        return "", err
    }

	return string(output), nil
}

func RunGitCommit(message string) error {
	args := []string{"commit","-m",message}
	cmd := exec.Command("git", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}