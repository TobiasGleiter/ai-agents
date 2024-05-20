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

	llamaResponse, err := ollama.Generate(llamaRequest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Response:", llamaResponse.Response)
}