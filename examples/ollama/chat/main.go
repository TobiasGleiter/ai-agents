package main

import (
	"fmt"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
) 

func main() {
	

	// messages = append(messages, ollama.ModelMessage{
	// 	Role: "user",
	// 	Content: "why is the sky blue?",
	// })

	wizardlm2_7b_model := ollama.OllamaModel{
		Model:  "wizardlm2:7b",
		Options: ollama.ModelOptions{
			Temperature: 0.7,
			NumCtx: 4096,
		},
		Stream: true,
	}

	ollamaClient := ollama.NewOllamaClient(wizardlm2_7b_model)

	prompt := "Hello, tell me a good joke!"
	_, err := ollamaClient.Chat(prompt)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// for _, message := range ollamaClient.Messages {
	// 	fmt.Println(string(message.Content))
	// }	

	// // Returns the final response after the stream is done.
	// _, err := ollama.Chat(llamaRequest)
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	return
	// }
}