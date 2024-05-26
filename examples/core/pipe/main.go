package main

import (
	"fmt"

	p "github.com/TobiasGleiter/ai-agents/pkg/model_io/output/string"
	//jp "github.com/TobiasGleiter/ai-agents/pkg/model_io/output/json"

	"github.com/TobiasGleiter/ai-agents/pkg/core/model"
	"github.com/TobiasGleiter/ai-agents/pkg/core/pipe"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
)

func main() {
	input := "hi, bye"

	llama3_8b_model := ollama.OllamaModel{
		Model:  "llama3:8b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream: false,
	}
	ollamaModelWrapper := model.NewOllamaModelWrapper(llama3_8b_model)

	//outputParser := &p.CommaSeparatedListOutputParser{}
	outputParser := &p.StringOutputParser{}
	//outputParser := &jp.JsonOutputParser{}

	inputModelOutputPipe := pipe.NewPipe(input, ollamaModelWrapper, outputParser)
	generatedOutput := inputModelOutputPipe.Invoke()
	fmt.Println(generatedOutput)
}
