package go_llama_agentic_system

import (
	"context"
	"os"
	"testing"

	"github.com/sashabaranov/go-openai"
)

func TestSystemPrompt(t *testing.T) {
	apiKey := os.Getenv("API_KEY")
	if os.Getenv("API_KEY") == "" {
		t.Fatal("environment variable API_KEY not set")
	}
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://api.fireworks.ai/inference/v1"
	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model: "accounts/fireworks/models/llama-v3p1-70b-instruct",
			Prompt: CreatePrompt(
				[]Message{
					{
						Role: SystemRole,
						Content: SystemPrompt(
							SystemPromptParams{
								IPythonEnabled: true,
								BuiltInTools:   []BuiltinTool{BraveSearch, WolframAlpha},
							},
						),
					},
					{
						Role:    UserRole,
						Content: `What is the biggest building in the world? Use brave search`,
					},
					{
						Role: AssistantRole,
					},
				},
			),
			Temperature: 0.5,
			MaxTokens:   4024,
			Stop:        []string{StartHeader},
		},
	)

	if err != nil {
		t.Errorf("completion error: %s", err)
		return
	}

	t.Log(resp.Choices[0].Text)
	t.Log(resp.Choices[0].FinishReason)
}
