package main 

import (
	"fmt"
	"encoding/json"
	"log"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
) 

type CompanyAndTicker struct {
    Company string `json:"company"`
    Ticker  string `json:"ticker"`
}

func main() {
	stockTickerNameOfCompany := "Google"

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
		Model:  "wizardlm2:7b",
		Messages: messages,
		Options: ollama.ModelOptions{
			Temperature: 0.8,
			NumCtx: 4096,
		},
		Stream: false,
		Format: "json",
	}

	// Returns the final response after the stream is done.
	limit := 4
	var response CompanyAndTicker
	
	for i := 0; i < limit; i++ {
		res, _ := ollama.Chat(llamaRequest)

		err = json.Unmarshal([]byte(res.Message.Content), &response)
		if err == nil {
			break
		}
		fmt.Println("Retrying", i)

	}

	if err != nil {
		log.Fatalf("Failed to decode JSON response after %d attempts: %v", limit, err)
	}

    fmt.Printf("Company: %s, Ticker: %s\n", response.Company, response.Ticker)
}