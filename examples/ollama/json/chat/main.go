package main 

import (
	"fmt"
	"log"
	"encoding/json"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
) 

type CompanyAndTicker struct {
    Company string `json:"company"`
    Ticker  string `json:"ticker"`
}

func main() {	
	wizardlm2_7b := ollama.OllamaModel{
		Model:  "wizardlm2:7b",
		Options: ollama.ModelOptions{
			Temperature: 0.7,
			NumCtx: 4096,
		},
		Stream: false,
		Format: "json",
		KeepAlive: -1,
	}
	
	ollamaClient := ollama.NewOllamaClient(wizardlm2_7b)

	systemPrompt := fmt.Sprintf(`
		You are a helpful AI assistant.
		The user will enter a company name and the assistant will output the response in JSON format like this:
		{"company": "Apple", "ticker": "AAPL"}`)

	
	var fewShotMessages []ollama.ModelMessage
	fewShotMessages = append(fewShotMessages, ollama.ModelMessage{
		Role: "user",
		Content: "Apple. Respond in JSON.", // Necessary to add "Respond in JSON" or there will be many whitespaces
	})
	fewShotMessages = append(fewShotMessages, ollama.ModelMessage{
		Role: "assistant",
		Content: `{"company": "Apple", "ticker": "AAPL"}`,
	})
	
	ollamaClient.SetSystemPrompt(systemPrompt)
	ollamaClient.SetMessages(fewShotMessages) // Optionally, depending on the accuracy of the generated output

	var response CompanyAndTicker
	res, err := ollamaClient.Chat("Google")
	
	err = json.Unmarshal([]byte(res.Message.Content), &response)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	fmt.Printf("\nDecoded JSON: Company: %s, Ticker: %s\n", response.Company, response.Ticker)
}