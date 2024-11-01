package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type promptContent struct {
	errorMsg string
	label    string
}

func PromptGetSelect(pc promptContent) string {
	items := []string{"yes", "no"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.Select{
			Label:        pc.label,
			Items:        items,
			HideSelected: true,
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("User input select prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}
