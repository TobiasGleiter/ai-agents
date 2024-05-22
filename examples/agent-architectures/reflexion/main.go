package main 

import (
	"fmt"
	"encoding/json"
	"log"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
	ChatColor "github.com/TobiasGleiter/ai-agents/internal/color"
)

type Tool struct {
	Name string `json:"name"`
}

type InitialResponse struct {
    Response string `json:"response"`
    Critque  string `json:"critque"`
	Tools []Tool `json:"tools"`
}

var functions = map[string]func(){
	"search_internet": searchInternet,
	"current_weather": getCurrentWeather,
	"save_file": saveFile,
}

func searchInternet() {
	ChatColor.PrintColor(ChatColor.Red, "Searching the internet...")
}

func getCurrentWeather() {
	ChatColor.PrintColor(ChatColor.Red, "Getting current weather...")
}

func saveFile() {
	ChatColor.PrintColor(ChatColor.Red, "Saving file...")
}

func main() {
	// Reflexion-Actor Architecture

	// 1. User input
	// 2. Initial response is generated along with self critque and suggested tool queries (needs a list of tools.)
	// 3. Use the tool(s)
	// 4. Pass the initial prompt and response to the Revisor LLM
	// 5. Iterate until the Revisor LLM returns true


	// 1. User input
	userInput := "Why is reflection useful in AI?"

	// 2. Initial Response LLM.
	// Output a JSON that can be handled accordingly.
	// Use few shot prompting to increase the quality of the prompt.
	initalResponseLlm := `
	{
		"response": "",
		"critque": "",
		"tools": [
			{
				"name": "tool_name"
			}
		]
	}`

	tools := []Tool{{Name: "search_internet"}, {Name: "current_weather"}, {Name: "save_file"}}

	var messages []ollama.ModelMessage
	messages = append(messages, ollama.ModelMessage{
        Role: "system",
        Content: fmt.Sprintf(`
			You are a helpful AI assistant.
			Generate an inital response along with self critque (how to improve the response) and select the right tools that helps solve the problem.
			Tools available: %s
			Respond in JSON format like this:
				%s`, tools, initalResponseLlm),
    })

	messages = append(messages, ollama.ModelMessage{
		Role: "user",
		Content: userInput,
	})

	llamaRequest := ollama.Model{
		Model:  "llama3:8b",
		Messages: messages,
		Options: ollama.ModelOptions{
			Temperature: 1,
			NumCtx: 4096,
		},
		Stream: false,
		Format: "json",
	}

	var response InitialResponse
	res, err := ollama.Chat(llamaRequest)

	err = json.Unmarshal([]byte(res.Message.Content), &response)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %s", err)
	}

	ChatColor.PrintColor(ChatColor.Yellow, "Initial Response: " + string(response.Response))
	ChatColor.PrintColor(ChatColor.Cyan, "Initial Critque: " + string(response.Critque))
	ChatColor.PrintColor(ChatColor.Green, "Tools: " + fmt.Sprintf("%s", response.Tools))

	// 3. Use the tool(s)
	for i := 0; i < len(response.Tools); i++ {
		if function, exists := functions[response.Tools[i].Name]; exists {
			function()
		} else {
			fmt.Printf("No function found. \n")
		}
	}


	// 4. Pass the initial prompt and response to the Revisor LLM
	var revisorMessages []ollama.ModelMessage
	revisorMessages = append(revisorMessages, ollama.ModelMessage{
        Role: "system",
        Content: fmt.Sprintf(`
			You are a revisor AI assistant.

			Use the following response and the critque to improve the response.
			If the response is really good, set Done to true.
			
			Respond in JSON format like this:
				%s`, initalResponseLlm),
    })

	limit := 2
	for j := 0; j < limit; j++ {
		revisorMessages = append(revisorMessages, ollama.ModelMessage{
			Role: "user",
			Content: fmt.Sprintf(`
				User Response: %s
				Critque: %s`, response.Response, response.Critque),
		})

		llamaRevisorRequest := ollama.Model{
			Model:  "llama3:8b",
			Messages: revisorMessages,
			Options: ollama.ModelOptions{
				Temperature: 1,
				NumCtx: 4096,
			},
			Stream: false,
			Format: "json",
		}

		res, err = ollama.Chat(llamaRevisorRequest)
		err = json.Unmarshal([]byte(res.Message.Content), &response)
		if err != nil {
			log.Fatalf("Failed to decode JSON: %s", err)
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Failed to encode JSON: %s", err)
		}

		revisorMessages = append(revisorMessages, ollama.ModelMessage{
			Role: "assistant",
			Content: string(jsonResponse),
		})

		ChatColor.PrintColor(ChatColor.Yellow, "Initial Response: " + string(response.Response))
		ChatColor.PrintColor(ChatColor.Cyan, "Initial Critque: " + string(response.Critque))
		ChatColor.PrintColor(ChatColor.Yellow, fmt.Sprintf("Revision: %v", j))
	}

	ChatColor.PrintColor(ChatColor.Green, "Final Response: " + response.Response)
}