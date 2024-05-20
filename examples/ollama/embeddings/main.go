package main 

import (
	"fmt"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
) 

func main() {
	llamaRequest := ollama.Model{
		Model:  "llama3:8b",
		Prompt: "Hi, how do you do?",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream: false,
	}

	llamaEmbeddingResponse, err := ollama.GenerateEmbeddings(llamaRequest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Response:", llamaEmbeddingResponse.Embedding)
}