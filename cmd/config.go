package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/maran1947/localite/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configurations",
	Long:  `Manage application configurations using key-value pairs.`,
	Run: func(cmd *cobra.Command, args []string) {
		listFlag, _ := cmd.Flags().GetBool("list")
		if listFlag {
			configData, err := config.LoadConfig()
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}

			configJsonData, err := json.MarshalIndent(configData, "", "  ")
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error marshalling config to JSON:", err)
				os.Exit(1)
			}

			fmt.Println(string(configJsonData))
			os.Exit(0)
		}
	},
}

var setCmd = &cobra.Command{
	Use:   "set KEY=VALUE",
	Short: "Set a configuration value",
	Long:  `Set a configuration value in the format KEY=VALUE.`,
	Run: func(cmd *cobra.Command, args []string) {
		listFlag, _ := cmd.Flags().GetBool("list")
		if listFlag {
			configData, err := config.LoadConfig()
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}

			fmt.Println(configData)
			os.Exit(0)
		}

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

var delCmd = &cobra.Command{
	Use:   "del KEY",
	Short: "del a key-value from configuration",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("Please provide the key to delete.")
			return
		}

		key := args[0]

		localiteConfig, err := config.LoadConfig()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("key: ", key)
		localiteConfig.DeleteConfigValue(key)

		err = config.UpdateConfig(*localiteConfig)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Configuration key '%s' has been deleted.\n", key)
		os.Exit(0)
	},
}

var getCmd = &cobra.Command{
	Use:   "get KEY",
	Short: "get a value of a given key from configuration",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("Please provide the key.")
			return
		}

		key := args[0]

		localiteConfig, err := config.LoadConfig()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
			os.Exit(1)
		}

		value, exists := localiteConfig.GetConfigValue(key)
		if !exists {
			fmt.Printf("Given key '%s' doesn not exists.\n", key)
			os.Exit(1)
		}

		fmt.Println(value)
		os.Exit(0)
	},
}

func init() {
	configCmd.PersistentFlags().BoolP("list", "l", false, "Displays the current configurations for the Localite tool")
	configCmd.AddCommand(setCmd)
	configCmd.AddCommand(delCmd)
	configCmd.AddCommand(getCmd)
}
