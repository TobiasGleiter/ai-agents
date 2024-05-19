package ollama

import (
	"net/http"
	"fmt"
	"encoding/json"
	"bytes"
)

type OllamaOptions struct {
	NumCtx int `json:"num_ctx,omitempty"`
}

type OllamaModel struct {
    Model    string `json:"model"`
    Prompt   string `json:"prompt"`
    Options  OllamaOptions
    Stream bool `json:"stream"`
	Format   string   `json:"format,omitempty"`
	KeepAlive int64 `json:"keepalive,omitempty"`
}

type OllamaResponse struct {
	Model              string        `json:"model"`
	CreatedAt          string        `json:"created_at"`
	Response           string        `json:"response"`
	Done               bool          `json:"done"`
	Context            []interface{} `json:"context"`
	TotalDuration      int64         `json:"total_duration"`
	LoadDuration       int64         `json:"load_duration"`
	PromptEvalCount    int           `json:"prompt_eval_count"`
	PromptEvalDuration int64         `json:"prompt_eval_duration"`
	EvalCount          int           `json:"eval_count"`
	EvalDuration       int64         `json:"eval_duration"`
}

const (
	Reset  = "\033[0m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func Generate(prompt string) string {
	client := &http.Client{}
	llmEndpoint := "http://localhost:11434/api/generate"
	
	request := OllamaModel{
        Model:  "llama3:8b",
        Prompt: prompt,
		Options: OllamaOptions{
            NumCtx: 4096,
        },
        Stream: false,
    }

	fmt.Println("\nPrompt: " + Yellow + prompt + Reset)

    requestBody, err := json.Marshal(request)
    if err != nil {
        fmt.Println("Error marshaling request:", err)
        return "error"
    }

	req, err := http.NewRequest("POST", llmEndpoint, bytes.NewReader(requestBody))
	if err != nil {
		fmt.Println("create request failed:", err)
		return "Error"
	}
	req.Header.Add("Content-Type", "application/json")
		

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("create request failed:", err)
		return "Error"
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var response OllamaResponse
	if err := decoder.Decode(&response); err != nil {	
		fmt.Println("error decoding response:", err)
	}

	fmt.Print(Green + response.Response + Reset)
	return response.Response
}