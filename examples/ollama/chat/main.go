package main

import (
	"fmt"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
) 

func main() {
	var messages []ollama.ModelMessage
	messages = append(messages, ollama.ModelMessage{
		Role: "user",
		Content: "why is the sky blue?",
	})

	llamaRequest := ollama.Model{
		Model:  "wizardlm2:7b",
		Messages: messages,
		Options: ollama.ModelOptions{
			Temperature: 0.8,
			NumCtx: 4096,
		},
		Stream: true,
	}

	// Returns the final response after the stream is done.
	_, err := ollama.Chat(llamaRequest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}