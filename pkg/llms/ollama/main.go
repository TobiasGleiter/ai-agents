package ollama

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	"github.com/pkg/errors"
)

type ModelOptions struct {
	NumCtx int `json:"num_ctx,omitempty"`
}

type Model struct {
    Model    string `json:"model"`
    Prompt   string `json:"prompt"`
    Options  ModelOptions
    Stream bool `json:"stream"`
	Format   string   `json:"format,omitempty"`
	KeepAlive int64 `json:"keepalive,omitempty"`
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

type EmbeddingResponse struct {
	Embedding []float64 `json:"embedding"`
}

const (
	llmGenerateEndpoint = "http://localhost:11434/api/generate"
	llmGenerateEmbeddingsEndpoint = "http://localhost:11434/api/embeddings"
)


func Generate(request interface{}) (Response, error) {
	client := &http.Client{
		Timeout: 240 * time.Second,
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return Response{}, errors.Wrap(err, "error marshaling request")
	}

	req, err := http.NewRequest("POST", llmGenerateEndpoint, bytes.NewReader(requestBody))
	if err != nil {
		return Response{}, errors.Wrap(err, "create request failed")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return Response{}, errors.Wrap(err, "HTTP request failed")
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return Response{}, errors.Wrap(err, "error decoding response")
	}

	return response, nil
}

func GenerateEmbeddings(request interface{}) (EmbeddingResponse, error) {
	client := &http.Client{
		Timeout: 240 * time.Second,
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return EmbeddingResponse{}, errors.Wrap(err, "error marshaling request")
	}

	req, err := http.NewRequest("POST", llmGenerateEmbeddingsEndpoint, bytes.NewReader(requestBody))
	if err != nil {
		return EmbeddingResponse{}, errors.Wrap(err, "create request failed")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return EmbeddingResponse{}, errors.Wrap(err, "HTTP request failed")
	}
	defer resp.Body.Close()

	var embeddingResponse EmbeddingResponse
	err = json.NewDecoder(resp.Body).Decode(&embeddingResponse)
	if err != nil {
		return EmbeddingResponse{}, errors.Wrap(err, "error decoding response")
	}

	return embeddingResponse, nil
}
