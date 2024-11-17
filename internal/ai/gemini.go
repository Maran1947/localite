package ai

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/maran1947/localite/internal/config"
	"google.golang.org/api/option"
)

func getPrompt(gitDiffData string, length int, allowPrefix bool) string {
	prefixDetails := `
Allow prefixes (docs, style, refactor, perf, test, build, ci, chore, revert, feat, fix).
1. docs: Changes related to documentation only.
2. style: Code formatting or non-functional changes (e.g., white-space).
3. refactor: Restructuring code without changing behavior.
4. perf: Performance improvements for faster code or reduced resources.
5. test: Adding or updating tests, no production code changes.
6. build: Changes affecting build systems or dependencies.
7. ci: Modifications to CI/CD configuration or scripts.
8. chore: Maintenance tasks, such as dependency updates.
9. revert: Undoing a previous commit due to issues.
10. feat: Implementing a new feature or enhancement.
11. fix: Bug fixes or patches for existing issues.
`
	if !allowPrefix {
		prefixDetails = "Don't use any prefixes including (docs, style, refactor, perf, test, build, ci, chore, revert, feat, fix)."
	}

	return fmt.Sprintf(`
Write a commit message summarizing the following git diff changes in a concise manner, limited to %d characters. Use clear and descriptive language, and avoid unnecessary details. Always write the commit message in the imperative mood, in lowercase letters, and do not include any explanations.
%s

Git diff:
%s

Commit message (within %d characters):`, length, prefixDetails, gitDiffData, length)
}

func GetResponse(gitDiffData string, length int, allowPrefix bool) (string, error) {
	ctx := context.Background()
	configData, err := config.LoadConfig()
	if err != nil {
		return "", err
	}

	geminiApiKey, isExists := configData.GetConfigValue("GEMINI_API_KEY")
	if !isExists {
		return "", fmt.Errorf("No Gemini API key exists. Please ensure the [GEMINI_API_KEY] is defined in your localite configuration")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(geminiApiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	prompt := getPrompt(gitDiffData, length, allowPrefix)

	response, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	if len(response.Candidates) == 0 || len(response.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response received from the model")
	}

	commitText := fmt.Sprint(response.Candidates[0].Content.Parts[0])

	return commitText, nil
}
