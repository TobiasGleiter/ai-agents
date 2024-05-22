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

func main() {
	// Reflexion-Actor Architecture

	// 1. User input
	// 2. Initial response is generated along with self critque and suggested tool queries (needs a list of tools.)
	// 3. Use the tool
	// 4. Pass the initial prompt and response to the Revisor LLM


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

	tools := []Tool{{Name: "search_internet"}, {Name: "get_current_weather"}, {Name: "save_file"}}

	var messages []ollama.ModelMessage
	messages = append(messages, ollama.ModelMessage{
        Role: "system",
        Content: fmt.Sprintf(`
			You are a helpful AI assistant.
			Generate an inital response along with self critque and select the right tools that helps solve the problem.
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

	// formattedJSON, err := json.MarshalIndent(chatResponse, "", "  ")
	// if err != nil {
	// 	log.Fatalf("Failed to format JSON: %s", err)
	// }

	ChatColor.PrintColor(ChatColor.Yellow, "Initial Response: " + string(response.Response))
	ChatColor.PrintColor(ChatColor.Cyan, "Initial Critque: " + string(response.Critque))
	ChatColor.PrintColor(ChatColor.Green, "Tools: " + fmt.Sprintf("%s", response.Tools))

	for i := 0; i < len(response.Tools); i++ {
		fmt.Printf(response.Tools[i].Name)
	}
}