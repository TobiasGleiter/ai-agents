package ollama

import (
	"fmt"
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	"io"
	"strings"

	"github.com/pkg/errors"
)

type ModelOptions struct {
	NumCtx int `json:"num_ctx"`
	Temperature float64 `json:"temperature"`
}

type ModelMessage struct {
	Role string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
	Images []string	`json:"images,omitempty"`
}

type Model struct {
    Model    string `json:"model"`
    Prompt   string `json:"prompt"`
	Messages []ModelMessage `json:"messages"`
    Options  ModelOptions
    Stream bool `json:"stream"`
	Format   string   `json:"format,omitempty"`
	KeepAlive int64 `json:"keepalive,omitempty"`
	Stop   []string `json:"stop"` // Not from Ollama
}

type Response struct {
	Model              string        `json:"model"`
	CreatedAt          string        `json:"created_at"`
	Response           string        `json:"response"`
	Message			   ModelMessage
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

type ChatResponse struct {
	Model				string		`json:"model"`
	CreatedAt			string		`json:"created_at"`
	Message 			ModelMessage
	Done				bool		`json:"done"`
	TotalDuration      int64         `json:"total_duration"`
	LoadDuration       int64         `json:"load_duration"`
	PromptEvalCount    int           `json:"prompt_eval_count"`
	PromptEvalDuration int64         `json:"prompt_eval_duration"`
	EvalCount          int           `json:"eval_count"`
	EvalDuration       int64         `json:"eval_duration"`
}

const (
	timeout = 240
	llmGenerateEndpoint = "http://localhost:11434/api/generate"
	llmGenerateEmbeddingsEndpoint = "http://localhost:11434/api/embeddings"
	llmChatEndpoint = "http://localhost:11434/api/chat"
)

type OllamaModel struct {
	Model    string `json:"model"`
    Options  ModelOptions
    Stream bool `json:"stream"`
	Format   string   `json:"format,omitempty"`
	KeepAlive int64 `json:"keepalive,omitempty"`
	Stop   []string `json:"stop"` // Not from Ollama
}

type OllamaGenerateRequest struct {
    Model    string `json:"model"`
    Prompt   string `json:"prompt"`
    Options  ModelOptions
    Stream bool `json:"stream"`
	Format   string   `json:"format,omitempty"`
	KeepAlive int64 `json:"keepalive,omitempty"`	
}

type OllamaChatRequest struct {
    Model    string `json:"model"`
	Messages []ModelMessage `json:"messages"`
    Options  ModelOptions
    Stream bool `json:"stream"`
	Format   string   `json:"format,omitempty"`
	KeepAlive int64 `json:"keepalive,omitempty"`	
}

type OllamaClient struct {
	Model OllamaModel
	Messages []ModelMessage
}

func NewOllamaClient(model OllamaModel) *OllamaClient {
	return &OllamaClient{Model: model}
}

func (oc *OllamaClient) Generate(prompt string) (Response, error) {
	client := &http.Client{
		Timeout: timeout * time.Second,
	}

	request := OllamaGenerateRequest{
		Model: oc.Model.Model,
		Prompt: prompt,
		Options: oc.Model.Options,
		Stream: oc.Model.Stream,
		Format: oc.Model.Format,
		KeepAlive: oc.Model.KeepAlive,
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

func (oc *OllamaClient) GenerateEmbeddings(prompt string) (EmbeddingResponse, error) {
	client := &http.Client{
		Timeout: timeout * time.Second,
	}

	request := OllamaGenerateRequest{
		Model: oc.Model.Model,
		Prompt: prompt,
		Options: oc.Model.Options,
		Stream: oc.Model.Stream,
		Format: oc.Model.Format,
		KeepAlive: oc.Model.KeepAlive,
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

func (oc *OllamaClient) Chat(prompt string) (ChatResponse, error) {
	client := &http.Client{
        Timeout: timeout * time.Second,
    }

	messages := append(oc.Messages, ModelMessage{
		Role: "user",
		Content: prompt,
	})

	fmt.Println(messages[0].Content)

	request := OllamaChatRequest{
		Model: oc.Model.Model,
		Messages: messages,
		Options: oc.Model.Options,
		Stream: oc.Model.Stream,
	}

	fmt.Println(request.Model)

    requestBody, err := json.Marshal(request)
    if err != nil {
        return ChatResponse{}, errors.Wrap(err, "error marshaling request")
    }

    req, err := http.NewRequest("POST", llmChatEndpoint, bytes.NewReader(requestBody))
    if err != nil {
        return ChatResponse{}, errors.Wrap(err, "create request failed")
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return ChatResponse{}, errors.Wrap(err, "HTTP request failed")
    }
	defer resp.Body.Close()

    
	decoder := json.NewDecoder(resp.Body)
	var chatResponse ChatResponse
	for {
		if err := decoder.Decode(&chatResponse); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return ChatResponse{}, errors.Wrap(err, "error decoding response")
		}
	}

	if chatResponse.Done {
		return chatResponse, nil
	}

	

    return chatResponse, nil	
}

func Generate(request Model) (Response, error) {
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

func GenerateEmbeddings(request Model) (EmbeddingResponse, error) {
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

func Chat(request Model) (ChatResponse, error) {
    client := &http.Client{
        Timeout: 240 * time.Second,
    }

    requestBody, err := json.Marshal(request)
    if err != nil {
        return ChatResponse{}, errors.Wrap(err, "error marshaling request")
    }

    req, err := http.NewRequest("POST", llmChatEndpoint, bytes.NewReader(requestBody))
    if err != nil {
        return ChatResponse{}, errors.Wrap(err, "create request failed")
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return ChatResponse{}, errors.Wrap(err, "HTTP request failed")
    }
	defer resp.Body.Close()

    
	decoder := json.NewDecoder(resp.Body)
	var chatResponse ChatResponse
	for {
		if err := decoder.Decode(&chatResponse); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return ChatResponse{}, errors.Wrap(err, "error decoding response")
		}
	}

	fmt.Println(chatResponse.Message.Content)

	if chatResponse.Done {
		return chatResponse, nil
	}


    return chatResponse, nil
}

func ChatReAct(request Model) (ChatResponse, error) {
	client := &http.Client{
		Timeout: 240 * time.Second,
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return ChatResponse{}, errors.Wrap(err, "error marshaling request")
	}

	req, err := http.NewRequest("POST", llmChatEndpoint, bytes.NewReader(requestBody))
	if err != nil {
		return ChatResponse{}, errors.Wrap(err, "create request failed")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return ChatResponse{}, errors.Wrap(err, "HTTP request failed")
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var chatResponse ChatResponse
	var finalResponse ChatResponse
	for {
		if err := decoder.Decode(&chatResponse); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return ChatResponse{}, errors.Wrap(err, "error decoding response")
		}
		finalResponse.Message.Content += chatResponse.Message.Content
		for _, stopSeq := range request.Stop {
			if strings.Contains(finalResponse.Message.Content, stopSeq) {
				finalResponse.Message.Content = strings.Split(finalResponse.Message.Content, stopSeq)[0]
				finalResponse.Done = true
				return finalResponse, nil
			}
		}
	}

	if chatResponse.Done {
		return finalResponse, nil
	}

	return finalResponse, nil
}


