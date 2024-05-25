package main 

import (
	"fmt"
	"encoding/json"
	"log"
	"os"

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

	userInput := "Why is reflection useful in AI?"

	llama3_8b_model := ollama.OllamaModel{
		Model:  "llama3:8b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream: false,
		Format: "json",
	}

	initialResponseLlm := ollama.NewOllamaClient(llama3_8b_model)
	revisorLlm := ollama.NewOllamaClient(llama3_8b_model)

	responseJsonFormat := `{
		"response": "",
		"critque": "",
		"tools": [{"name": "tool_name"}]
	}`

	tools := []Tool{{Name: "search_internet"}, {Name: "current_weather"}}
	systemPrompt := fmt.Sprintf(`
		You are a helpful AI assistant.
		Generate an inital response along with self critque (how to improve the response) and select the right tools that helps solve the problem.
		Tools available: %s
		Respond in JSON format like this: %s
		`, responseJsonFormat, tools)

	var fewShotMessages []ollama.ModelMessage
	fewShotMessages = append(fewShotMessages, ollama.ModelMessage{
		Role: "user",
		Content: "When is the next US election?", // Necessary to add "Respond in JSON" or there will be many whitespaces
	})
	fewShotMessages = append(fewShotMessages, ollama.ModelMessage{
		Role: "assistant",
		Content: `
			"response": "",
			"critque": "",
			"tools": [{"name": "search_internet"}]
		}`,
	})

	initialResponseLlm.SetSystemPrompt(systemPrompt)
	initialResponseLlm.SetMessages(fewShotMessages)
	res, err := initialResponseLlm.Chat(userInput)
	
	var response InitialResponse
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

	revisorSystemPrompt := fmt.Sprintf(`
	You are a revisor AI assistant.
	Generate a new better response along as a new critque to improve the response. 
	Respond in JSON format like this: %s
	`, responseJsonFormat)

	revisorLlm.SetSystemPrompt(revisorSystemPrompt)

	var finalResponse ollama.ChatResponse
	var revisorResponse InitialResponse
	var prompt string
	prompt = fmt.Sprintf("Response: %s; Critque: %s;", string(response.Response), string(response.Critque))

	n := 2
	for i := 0; i < n; i++ {
		finalResponse, _ = revisorLlm.Chat(prompt)

		err = json.Unmarshal([]byte(finalResponse.Message.Content), &revisorResponse)
		if err != nil {
			prompt = "Try again please."
			return
		}

		ChatColor.PrintColor(ChatColor.Cyan, revisorResponse.Critque)

		prompt = fmt.Sprintf("Response: %s; Critque: %s;", string(response.Response), string(response.Critque))
	}

	err = json.Unmarshal([]byte(finalResponse.Message.Content), &revisorResponse)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %s", err)
	}

	ChatColor.PrintColor(ChatColor.Green, "Final Response: " + revisorResponse.Response)
}


func writeMessagesToMarkdown(messages []ollama.ModelMessage, filename string) error {
    // Create or open the markdown file
    file, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("failed to create file: %w", err)
    }
    defer file.Close()

    // Write the messages to the markdown file
    for _, message := range messages {
        var header string
        if message.Role == "user" {
            header = "### User\n"
        } else if message.Role == "assistant" {
            header = "### Assistant\n"
        } else {
            header = "### " + message.Role + "\n"
        }

        _, err := file.WriteString(header + "\n" + message.Content + "\n\n")
        if err != nil {
            return fmt.Errorf("failed to write to file: %w", err)
        }
    }

    return nil
}