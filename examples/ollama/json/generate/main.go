package main 

import (
	"fmt"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
) 

func main() {
	wizardlm2_7b := ollama.OllamaModel{
		Model:  "wizardlm2:7b",
		Options: ollama.ModelOptions{
			Temperature: 0.7,
			NumCtx: 4096,
		},
		Stream: false, // If generate then this need to be false!
		Format: "json",
	}
	
	ollamaClient := ollama.NewOllamaClient(wizardlm2_7b)

	prompt := `
		You are a helpful AI assistant.
		The User will ask a question and the assistant will output the response in JSON format like this:
		{"answer": ""}
		
		What color is the sky at different times of the day? Respond using JSON.`

	res, err := ollamaClient.Generate(prompt)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Response:", res.Response)
}