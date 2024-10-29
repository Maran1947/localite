package utils

import (
	"os/exec"
	"strings"
)

func GetGitDiff(flags string) (string, error) {
	cmdArgs := []string{"diff","."}
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