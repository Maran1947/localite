package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/maran1947/localite/internal/ai"
	"github.com/maran1947/localite/internal/utils"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
    Use:   "generate",
    Short: "Generate commit message",
    Long:  `Generate commit message for the current git changes.`,
    Run: func(cmd *cobra.Command, args []string) {
        length, _ := cmd.Flags().GetInt("length")
        if length <= 0 {
            fmt.Println("Length must be a positive integer.")
            return
        }
        gitDiffData, err:= utils.GetGitDiff(strings.Join(args, " "))
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error: %v\n", err)
            os.Exit(1)
        }
        
        if len(gitDiffData) == 0 {
            fmt.Println("There is not git changes found!")
            return
        }

        ai.GetResponse(gitDiffData, length)
    },
}

func init() {
    generateCmd.Flags().IntP("length", "l", 0, "Length of the random text to generate")
}
