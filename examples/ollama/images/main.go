package main

import (
	"fmt"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
	ImageTool "github.com/TobiasGleiter/ai-agents/pkg/tools/image"
) 

func main() {
	var messages []ollama.ModelMessage
	var images []string

	ImageTool.LoadImageFromPath("./objectdetection.jpg")
	base64image := ImageTool.EncodeImageToBase64()
	images = append(images, base64image)

	messages = append(messages, ollama.ModelMessage{
		Role: "user",
		Content: "What is in this picture?",
		Images: images,
	})

	llavaRequest := ollama.Model{
		Model:  "llava:7b",
		Messages: messages,
		Options: ollama.ModelOptions{ NumCtx: 4096, },
		Stream: false, // Alternativly use true for streaming
	}

	// Returns the final response after the stream is done.
	_, err := ollama.Chat(llavaRequest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}