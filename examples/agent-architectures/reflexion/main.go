package main 

import (
	"fmt"
	// "encoding/json"
	// "log"
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

	tools := []Tool{{Name: "search_internet"}, {Name: "current_weather"}, {Name: "save_file"}}
	systemPrompt := fmt.Sprintf(`
		You are a helpful AI assistant.
		Generate an inital response along with self critque (how to improve the response) and select the right tools that helps solve the problem.
		Tools available: %s
		Respond in JSON format like this:
		{
			"response": "",
			"critque": "",
			"tools": [
				{
					"name": "tool_name"
				}
			]
		}`, tools)

	var fewShotMessages []ollama.ModelMessage
	fewShotMessages = append(fewShotMessages, ollama.ModelMessage{
		Role: "user",
		Content: "When is the next US president election?", // Necessary to add "Respond in JSON" or there will be many whitespaces
	})
	fewShotMessages = append(fewShotMessages, ollama.ModelMessage{
		Role: "assistant",
		Content: `
			"response": "I dont have the right answer, I need to search the internet.",
			"critque": "",
			"tools": [
				{
					"name": "search_internet",
				}
			]
		}`,
	})

	initialResponseLlm.SetSystemPrompt(systemPrompt)
	initialResponseLlm.SetMessages(fewShotMessages)
	initialResponseLlm.Chat(userInput)

	// var messages []ollama.ModelMessage
	// messages = append(messages, ollama.ModelMessage{
    //     Role: "system",
    //     Content: fmt.Sprintf(`
	// 		You are a helpful AI assistant.
	// 		Generate an inital response along with self critque (how to improve the response) and select the right tools that helps solve the problem.
	// 		Tools available: %s
	// 		Respond in JSON format like this:
	// 			%s`, tools, initalResponseLlm),
    // })

	// messages = append(messages, ollama.ModelMessage{
	// 	Role: "user",
	// 	Content: userInput,
	// })

	// llamaRequest := ollama.Model{
	// 	Model:  "llama3:8b",
	// 	Messages: messages,
	// 	Options: ollama.ModelOptions{
	// 		Temperature: 1,
	// 		NumCtx: 4096,
	// 	},
	// 	Stream: false,
	// 	Format: "json",
	// }

	// var response InitialResponse
	// res, err := ollama.Chat(llamaRequest)

	// err = json.Unmarshal([]byte(res.Message.Content), &response)
	// if err != nil {
	// 	log.Fatalf("Failed to decode JSON: %s", err)
	// }

	// ChatColor.PrintColor(ChatColor.Yellow, "Initial Response: " + string(response.Response))
	// ChatColor.PrintColor(ChatColor.Cyan, "Initial Critque: " + string(response.Critque))
	// ChatColor.PrintColor(ChatColor.Green, "Tools: " + fmt.Sprintf("%s", response.Tools))

	// // 3. Use the tool(s)
	// for i := 0; i < len(response.Tools); i++ {
	// 	if function, exists := functions[response.Tools[i].Name]; exists {
	// 		function()
	// 	} else {
	// 		fmt.Printf("No function found. \n")
	// 	}
	// }


	// var revisorMessages []ollama.ModelMessage
	// // 4. Pass the initial prompt and response to the Revisor LLM
	// revisorMessages = append(revisorMessages, ollama.ModelMessage{
    //     Role: "system",
    //     Content: fmt.Sprintf(`
	// 		You are a revisor AI assistant.

	// 		Generate a new better response along as a new critque to improve the response.
			
	// 		Respond in JSON format like this:
	// 			%s`, initalResponseLlm),
    // })

	// limit := 4
	// for j := 0; j < limit; j++ {
	// 	revisorMessages = append(revisorMessages, ollama.ModelMessage{
	// 		Role: "user",
	// 		Content: fmt.Sprintf(`
	// 			Generate a new better response along as a new critque to improve the response.

	// 			Response: %s
	// 			Critque: %s

	// 			Respond in JSON.`, response.Response, response.Critque),
	// 	})

	// 	llamaRevisorRequest := ollama.Model{
	// 		Model:  "llama3:8b",
	// 		Messages: revisorMessages,
	// 		Options: ollama.ModelOptions{
	// 			Temperature: 0.8,
	// 			NumCtx: 4096,
	// 		},
	// 		Stream: false,
	// 		Format: "json",
	// 	}

	// 	res, err = ollama.Chat(llamaRevisorRequest)
	// 	err = json.Unmarshal([]byte(res.Message.Content), &response)
	// 	if err != nil {
	// 		log.Fatalf("Failed to decode JSON: %s", err)
	// 	}

	// 	jsonResponse, err := json.Marshal(response)
	// 	if err != nil {
	// 		log.Fatalf("Failed to encode JSON: %s", err)
	// 	}

	// 	revisorMessages = append(revisorMessages, ollama.ModelMessage{
	// 		Role: "assistant",
	// 		Content: string(jsonResponse),
	// 	})

	// 	ChatColor.PrintColor(ChatColor.Yellow, "Initial Response: " + string(response.Response))
	// 	ChatColor.PrintColor(ChatColor.Cyan, "Initial Critque: " + string(response.Critque))
	// 	ChatColor.PrintColor(ChatColor.Yellow, fmt.Sprintf("Revision: %v", j))
	// }

	// ChatColor.PrintColor(ChatColor.Green, "Final Response: " + response.Response)

	// err = writeMessagesToMarkdown(revisorMessages, "messages.md")
    // if err != nil {
    //     log.Fatalf("Error writing to markdown file: %s", err)
    // }
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