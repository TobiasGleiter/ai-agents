package main

import (
	"fmt"

	"github.com/TobiasGleiter/ai-agents/pkg/pipe"
	pt "github.com/TobiasGleiter/ai-agents/pkg/model_io/output/string"
	"github.com/TobiasGleiter/ai-agents/pkg/core/model"

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

	outputParser := &pt.CommaSeparatedListOutputParser{}

	inputModelOutputPipe := pipe.NewPipe(input, ollamaModelWrapper, outputParser)
	generatedOutput := inputModelOutputPipe.Invoke()
	fmt.Println(generatedOutput)
}
