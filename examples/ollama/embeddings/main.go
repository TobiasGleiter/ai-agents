package main 

import (
	"fmt"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
) 

func main() {
	llama3_8b_model := ollama.OllamaModel{
		Model:  "llama3:8b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream: false,
	}

	ollamaClient := ollama.NewOllamaClient(llama3_8b_model)


	prompt := "Hi, how do you do?"
	response, err := ollamaClient.GenerateEmbeddings(prompt)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Response:", response.Embedding)
}