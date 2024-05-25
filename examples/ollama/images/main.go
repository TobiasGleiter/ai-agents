package main

import (
	//"fmt"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
	ImageTool "github.com/TobiasGleiter/ai-agents/pkg/tools/image"
) 

func main() {
	var images []string

	imageTool := ImageTool.NewImageTool()

	imageTool.LoadImageFromPath("./objectdetection.jpg")
	base64image := imageTool.EncodeImageToBase64()
	images = append(images, base64image)
	

	llava_7b := ollama.OllamaModel{
		Model:  "llava:7b",
		Options: ollama.ModelOptions{ NumCtx: 4096, },
		Stream: true,
	}

	ollamaClient := ollama.NewOllamaClient(llava_7b)

	prompt := "What is in this picture?"
	ollamaClient.MultimodalChat(prompt, images)
	ollamaClient.Chat("How is this person?") // Ask follow up questions about the image

}