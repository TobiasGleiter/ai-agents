package main 

import (
	"fmt"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
	ChatColor "github.com/TobiasGleiter/ai-agents/internal/color"
)

func main() {
	// Reflexion-Actor Architecture

	// 1. User input
	// 2. Initial response is generated along wiht self critque and suggested tool queries

	// 1. User input
	userInput := "Why is reflection useful in AI?"
	prompt := userInput + " Generate an inital response along with self critque and select a tool that helps solve the problem"

	userRequest := ollama.Model{
		Model:  "wizardlm2:7b",
		Prompt: prompt,
		Options: ollama.ModelOptions{NumCtx: 4096},
	}

	// 2. Initial Response LLM
	initialResponse, err := ollama.Generate(userRequest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	ChatColor.PrintColor(ChatColor.Green, "Final Response: " + initialResponse.Response)
}