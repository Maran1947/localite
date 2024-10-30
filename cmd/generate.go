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
		allowPrefix, _ := cmd.Flags().GetBool("prefix")

		if length <= 0 {
			fmt.Println("Length must be a positive integer.")
			return
		}

		gitDiffData, err := utils.RunGitDiff(strings.Join(args, " "))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		if len(gitDiffData) == 0 {
			fmt.Println("There is not git changes found!")
			return
		}

		commitText, err := ai.GetResponse(gitDiffData, length, allowPrefix)
		fmt.Println("----------: Generated commit message :----------")
		fmt.Print(commitText)
		fmt.Println("------------------------------------------------")

		if err != nil {
			fmt.Println("Error occurred in generating commit message: ", err)
			return
		}

		commitRequest := promptContent{
			"No commit message found.",
			"Would you like to commit this message?",
		}

		commitResponse := PromptGetSelect(commitRequest)

		switch commitResponse {
		case "yes":
			if err := utils.RunGitCommit(commitText); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to commit: %v\n", err)
				os.Exit(1)
			}
			os.Exit(0)
		case "no":
			os.Exit(0)
		default:
			fmt.Println("unexpected error")
			os.Exit(1)
		}
	},
}

func init() {
	generateCmd.Flags().IntP("length", "l", 0, "Length of the random text to generate")
	generateCmd.Flags().BoolP("prefix", "p", false, "Allow conventional prefix to beginning of a commit message")
}
