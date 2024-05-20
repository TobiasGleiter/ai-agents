package main 

import (
	"fmt"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
	ChatColor "github.com/TobiasGleiter/ai-agents/internal/color"
)

func main() {
	// Basic-Reflection Agent Architecture

	// 1. User input (first prompt)
	// 2. Inital response (first LLM) from the user input <- yellow color
	// 3. Response fed into a second LLM 
	// 4. Generates critiques and ideas for improvements <- blue color
	// 5. Response fed into first LLM
	// Notes: repeats n times (uncontrolled)

	var prompt string
	var firstResponse ollama.Response

	initPrompt := "Hello, how do you do?"
	prompt = initPrompt

	n := 4
	for i := 0; i < n; i++ {
		userRequest := ollama.Model{
			Model:  "llama3:8b",
			Prompt: prompt,
			Options: ollama.ModelOptions{NumCtx: 4096},
		}

		// Generator LLM
		var err error
		firstResponse, err = ollama.Generate(userRequest)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		ChatColor.PrintColor(ChatColor.Yellow, firstResponse.Response)
		reflectionPrompt := firstResponse.Response + " How to improve the answer to this response: " + initPrompt

		reflectionRequest := ollama.Model{
			Model:  "llama3:8b",
			Prompt: reflectionPrompt ,
			Options: ollama.ModelOptions{NumCtx: 4096},
		}

		// Reflection LLM
		basicReflectionResponse, err := ollama.Generate(reflectionRequest)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		ChatColor.PrintColor(ChatColor.Blue, basicReflectionResponse.Response)
		prompt = basicReflectionResponse.Response + " Improve your reponse with these tips. Output only the improved response in a sentence."
	}

	ChatColor.PrintColor(ChatColor.Green, "Final Response: " + firstResponse.Response)
}