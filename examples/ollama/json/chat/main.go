package main 

import (
	"fmt"
	"encoding/json"
	"log"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
) 

func main() {
	stockTickerNameOfCompany := "Microsoft Corporation"

	companyAndTickerSchema := map[string]map[string]string{
        "company": {
            "type":        "string",
            "description": "Name of the company",
        },
        "ticker": {
            "type":        "string",
            "description": "Ticker symbol of the company",
        },
    }

	companyAndTickerSchemaJSON, err := json.Marshal(companyAndTickerSchema)
    if err != nil {
        log.Fatalf("Error marshalling schema: %v", err)
    }


	var messages []ollama.ModelMessage
	messages = append(messages, ollama.ModelMessage{
        Role: "system",
        Content: fmt.Sprintf(`You are a helpful AI assistant.
				The user will enter a company name and the assistant will output the response in JSON format like this:
				%s`, string(companyAndTickerSchemaJSON)),
    })

	messages = append(messages, ollama.ModelMessage{
		Role: "user",
		Content: "Apple",
	})

	messages = append(messages, ollama.ModelMessage{
		Role: "assistant",
		Content: `{"company": "Apple", "ticker": "AAPL"}`,
	})

	messages = append(messages, ollama.ModelMessage{
		Role: "user",
		Content: stockTickerNameOfCompany,
	})

	llamaRequest := ollama.Model{
		Model:  "llama3:8b",
		Messages: messages,
		Options: ollama.ModelOptions{
			Temperature: 0.8,
			NumCtx: 4096,
		},
		Stream: false,
		Format: "json",
	}

	// Returns the final response after the stream is done.
	finalResponse, err := ollama.Chat(llamaRequest)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(finalResponse.Response)
}