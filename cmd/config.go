package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/maran1947/localite/internal/config"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configurations",
	Long:  `Manage application configurations using key-value pairs.`,
}

var setCmd = &cobra.Command{
	Use:   "set KEY=VALUE",
	Short: "Set a configuration value",
	Long:  `Set a configuration value in the format KEY=VALUE.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the key-value pair in the format: KEY=VALUE")
			return
		}
		keyValuePair := strings.SplitN(args[0], "=", 2)
		if len(keyValuePair) != 2 {
			fmt.Println("Please provide the key-value pair in the format: KEY=VALUE")
			return
		}
		key, value := keyValuePair[0], keyValuePair[1]
		err := config.SaveConfig(key, value)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	configCmd.AddCommand(setCmd)
}
