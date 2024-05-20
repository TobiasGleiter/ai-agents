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
		Model:  "llama3:8b",
		Messages: messages,
		Stream: true,
	}

	// Returns the final response after the stream is done.
	_, err := ollama.Chat(llamaRequest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}