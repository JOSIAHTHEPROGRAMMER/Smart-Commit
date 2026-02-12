package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/JOSIAHTHEPROGRAMMER/Smart-Commit/internal/config"
	"google.golang.org/genai"
)

const systemPrompt = `You are an expert at writing git commit messages following the Conventional Commits specification.

Your task is to analyze git diffs and generate high-quality commit messages.

STRICT RULES:
1. Format: type(scope): title
2. Title MUST be under 72 characters
3. Title MUST be in imperative mood (e.g., "add feature" not "added feature")
4. Body MUST have 2-4 bullet points explaining what changed
5. Valid types ONLY: feat, fix, docs, style, refactor, test, chore, perf, ci, build, revert
6. Scope is optional but preferred (e.g., api, ui, auth, db)
7. Be concise and technical
8. Focus on WHAT changed and WHY, not HOW

OUTPUT FORMAT:
type(scope): short imperative title under 72 chars

- First change explained
- Second change explained
- Third change explained (if applicable)

DO NOT include any explanations, markdown formatting, or extra text. Return ONLY the commit message.`

const systemPromptShort = `Generate a concise conventional commit message.

Rules:
- Format: type(scope): title
- Title under 72 chars
- One line body summary only
- Valid types: feat, fix, docs, style, refactor, test, chore, perf, ci, build, revert

Return ONLY the commit message.`

// GenerateCommitMessage uses Gemini to generate a commit message from a diff
func GenerateCommitMessage(diff string, cfg *config.Config) (string, error) {
	apiKey, err := config.GetGeminiAPIKey()
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create Gemini client: %v", err)
	}

	// Choose prompt based on style
	selectedPrompt := systemPrompt
	if cfg.Style == "short" {
		selectedPrompt = systemPromptShort
	}

	prompt := fmt.Sprintf("%s\n\nAnalyze this git diff and generate a conventional commit message:\n\n%s", selectedPrompt, diff)

	// Use model from config
	model := cfg.Model
	if model == "" {
		model = "gemini-3-flash-preview"
	}

	result, err := client.Models.GenerateContent(
		ctx,
		model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %v", err)
	}

	message := result.Text()
	message = strings.TrimSpace(message)

	// Fallback if response is empty or too short k
	if message == "" || len(message) < 10 {
		return "chore: update files\n\n- Update project files\n- Apply changes from diff", nil
	}

	return message, nil
}
