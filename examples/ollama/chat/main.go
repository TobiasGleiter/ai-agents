package main

import (
	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
) 

func main() {
	wizardlm2_7b_model := ollama.OllamaModel{
		Model:  "wizardlm2:7b",
		Options: ollama.ModelOptions{
			Temperature: 0.7,
			NumCtx: 4096,
		},
		Stream: true,
	}

	ollamaClient := ollama.NewOllamaClient(wizardlm2_7b_model)

	ollamaClient.Chat("Hello, tell me a good joke!")
	ollamaClient.Chat("Explain this joke!")
	ollamaClient.Chat("Summarize the joke in a bullet list!")
}