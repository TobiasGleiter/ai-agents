package main 

import (
	"fmt"
	"encoding/json"
	"log"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
) 


type Parameters struct {
	Properties map[string]interface{} `json:"properties"`
}

type FunctionCalling struct {
    Name       string     `json:"name"`
    Parameters Parameters `json:"parameters"`
}

type FunctionCallResponse interface {
	Process() string
}

type WeatherResponse struct {
	Location string
	Format   string
	Temperature string
}

func (wr WeatherResponse) Process() string {
	return fmt.Sprintf("\nCurrent weather for %s: Temperature is %s in %s format", wr.Location, wr.Temperature, wr.Format)
}

func getCurrentWeather(params Parameters) FunctionCallResponse {
	format, _ := params.Properties["format"].(string)
	location, _ := params.Properties["location"].(string)
	temperature := "68" // Fetch the weather temperature

	return WeatherResponse{Location: location, Format: format, Temperature: temperature}
}

var functions = map[string]func(Parameters) FunctionCallResponse{
	"get_current_weather": getCurrentWeather,
}

func main() {
	prompt := "What's the weather like in Detroit?"

	// Similar to the OpenAI format: https://platform.openai.com/docs/guides/function-calling
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
		Content: "What's the weather like in Berlin?",
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

	var response FunctionCalling
	res, _ := ollamaClient.Chat(prompt)

	err := json.Unmarshal([]byte(res.Message.Content), &response)
	if err != nil {
		log.Fatalf("Failed to decode JSON: %s", err)
	}

	if function, exists := functions[response.Name]; exists {
		result := function(response.Parameters)
		fmt.Println(result.Process())
	} else {
		fmt.Printf("No function found for name %s\n", response.Name)
	}
}