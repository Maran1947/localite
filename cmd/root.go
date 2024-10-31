package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const latestVersion = "v1.0.0"

var rootCmd = &cobra.Command{
	Use:   "localite",
	Short: "Localite CLI simplifies secret management and generates meaningful commit messages.",
	Long:  `A powerful CLI tool for handling local development functionalities, including config management and AI-powered commit generation.`,
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			fmt.Print(latestVersion)
			os.Exit(0)
		}
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Display the latest version of the Localite CLI")
}
