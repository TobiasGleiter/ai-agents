package main 

import (
	"fmt"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
) 

func main() {
	llamaRequest := ollama.Model{
		Model:  "llama3:8b",
		Prompt: "What color is the sky at different times of the day? Respond using JSON",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Format: "json",
		Stream: false,
	}

	llamaEmbeddingResponse, err := ollama.Generate(llamaRequest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Response:", llamaEmbeddingResponse.Response)
}