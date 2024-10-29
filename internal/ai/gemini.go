package ai

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/maran1947/localite/internal/config"
	"google.golang.org/api/option"
)

func getPrompt(gitDiffData string, length int) string {
	return fmt.Sprintf(`Write a commit message summarizing the following git diff changes in a concise manner, limited to %d characters. Aim to capture the key updates, improvements, or fixes reflected in the diff. Use clear and descriptive language, and avoid unnecessary details.

	Git diff:
	%s
	
	Commit message (within %d characters):`, length, gitDiffData, length)
}

func GetResponse(gitDiffData string, length int) (string, error) {
	ctx := context.Background()
	configData, err:= config.LoadConfig()
	if err != nil {
		return "", err
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(configData.GeminiApiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	prompt := getPrompt(gitDiffData, length)

    response, err := model.GenerateContent(ctx, genai.Text(prompt))
    if err != nil {
        log.Fatal(err)
    }

	if len(response.Candidates) == 0 || len(response.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response received from the model")
	}

	fmt.Println("Generated commit message:")
	fmt.Println(response.Candidates[0].Content.Parts[0])

	return "", nil
} 