package model

import (
	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
)

type OllamaModelWrapper struct {
	client *ollama.OllamaClient
}

func NewOllamaModelWrapper(model ollama.OllamaModel) *OllamaModelWrapper {
	client := ollama.NewOllamaClient(model)
	return &OllamaModelWrapper{client: client}
}

func (o *OllamaModelWrapper) Process(input string) string {
	response, err := o.client.Generate(input)
	if err != nil {
		return ""
	}
	return response.Response
}