package main

import (
	"fmt"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
) 

func main() {
	// Define model and options
	llama3_8b_model := ollama.OllamaModel{
		Model:  "llama3:8b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream: false,
	}

	// Initialize a new ollamaClient instance
	ollamaClient := ollama.NewOllamaClient(llama3_8b_model)

	// Send a prompt to the model
	prompt := "Hi, how do you do?"
	llamaResponse, err := ollamaClient.Generate(prompt)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Response:", llamaResponse.Response)

}