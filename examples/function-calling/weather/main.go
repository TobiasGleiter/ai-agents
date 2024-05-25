package main 

import (
	"fmt"
	//"encoding/json"
	//"log"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
	ChatColor "github.com/TobiasGleiter/ai-agents/internal/color"
) 

type Parameters struct {
	Properties map[string]interface{} `json:"properties"`
}

type FunctionCalling struct {
    Name string `json:"name"`
    Parameters  Parameters `json:"parameters"`
}

var functions = map[string]func(Parameters){
	"get_current_weather": getCurrentWeather,
}

func getCurrentWeather(params Parameters) {
	format, _ := params.Properties["format"].(string)
	location, _ := params.Properties["location"].(string)
	ChatColor.PrintColor(ChatColor.Cyan, "Getting current weather for " + location + " in " +  format)
}

func main() {
	prompt := "What's the weather like in Detroit?"

	// Similar to the OpenAI format
	weatherTool := `
	{
        "name": "get_current_weather",
        "description": "Get the current weather",
        "parameters": {
            "type": "object",
            "properties": {
                "location": {
                    "type": "string",
                    "description": "The city and state, e.g. San Francisco, CA"
                },
                "format": {
                    "type": "string",
                    "enum": [
                        "celsius",
                        "fahrenheit"
                    ],
                    "description": "The temperature unit to use. Infer this from the users location."
                }
            }
        }
    }`

	llama3_8b := ollama.OllamaModel{
		Model:  "llama3:8b",
		Options: ollama.ModelOptions{
			Temperature: 0.7,
			NumCtx: 4096,
		},
		Stream: false,
		Format: "json",
	}

	ollamaClient := ollama.NewOllamaClient(llama3_8b)

	systemPrompt := fmt.Sprintf(`
		You are a helpful AI assistant.
		Respond in JSON format like this:
		%s`, weatherTool)

	var fewShotMessages []ollama.ModelMessage
	fewShotMessages = append(fewShotMessages, ollama.ModelMessage{
		Role: "user",
		Content: "What's the weather like in Berlin?", // Necessary to add "Respond in JSON" or there will be many whitespaces
	})
	fewShotMessages = append(fewShotMessages, ollama.ModelMessage{
		Role: "assistant",
		Content: `{
			"name": "get_current_weather",
			"parameters": {
			 	"properties": {
			 		"location": "berlin",
			 		"format": "celsius"
			 	}
			}`,
	})

	ollamaClient.SetSystemPrompt(systemPrompt)
	ollamaClient.SetMessages(fewShotMessages)
	ollamaClient.Chat(prompt)




	// fewShot := `
	// {
	// 	"name": "get_current_weather",
	// 	"parameters": {
	// 		"properties": {
	// 			"location": "berlin",
	// 			"format": "celsius"
	// 		}
	// 	}
	// }`

	// var messages []ollama.ModelMessage
	// messages = append(messages, ollama.ModelMessage{
    //     Role: "system",
    //     Content: fmt.Sprintf(`
	// 		You are a helpful AI assistant.
	// 		Respond in JSON format like this:
	// 			%s`, weatherTool),
    // })

	// messages = append(messages, ollama.ModelMessage{
	// 	Role: "user",
	// 	Content: "What's the weather like in Berlin and Stuttgart?",
	// })

	// messages = append(messages, ollama.ModelMessage{
	// 	Role: "assistant",
	// 	Content: fewShot,
	// })

	// messages = append(messages, ollama.ModelMessage{
	// 	Role: "user",
	// 	Content: prompt,
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

	// var response FunctionCalling
	// res, err := ollama.Chat(llamaRequest)

	// err = json.Unmarshal([]byte(res.Message.Content), &response)
	// if err != nil {
	// 	log.Fatalf("Failed to decode JSON: %s", err)
	// }

	// if function, exists := functions[response.Name]; exists {
	// 	function(response.Parameters)
	// } else {
	// 	fmt.Printf("No function found for name %s\n", response.Name)
	// }
}