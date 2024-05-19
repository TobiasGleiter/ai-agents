package llama3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Options struct {
	NumCtx int `json:"num_ctx,omitempty"`
}

type Model struct {
	Model     string  `json:"model"`
	Prompt    string  `json:"prompt"`
	Options   Options `json:"options"`
	Stream    bool    `json:"stream"`
	Format    string  `json:"format,omitempty"`
	KeepAlive int64   `json:"keepalive,omitempty"`
}

type Response struct {
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

type Generator struct {
	Client     *http.Client
	Endpoint   string
	DefaultModel string
	DefaultOptions Options
}

// NewGenerator creates a new Generator with the given HTTP client, endpoint, and default model/options.
func NewGenerator(client *http.Client, endpoint string, defaultModel string, defaultOptions Options) *Generator {
	return &Generator{
		Client:     client,
		Endpoint:   endpoint,
		DefaultModel: defaultModel,
		DefaultOptions: defaultOptions,
	}
}

// Generate sends a prompt to the LLM API and returns the response.
func (g *Generator) Generate(prompt string) (string, error) {
	request := Model{
		Model:  g.DefaultModel,
		Prompt: prompt,
		Options: g.DefaultOptions,
		Stream: false,
	}

	fmt.Println("\nPrompt: " + Yellow + prompt + Reset)

	requestBody, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	req, err := http.NewRequest("POST", g.Endpoint, bytes.NewReader(requestBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := g.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	fmt.Print(Green + response.Response + Reset)
	return response.Response, nil
}
