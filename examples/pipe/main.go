package main

import (


	"github.com/TobiasGleiter/ai-agents/pkg/pipe"
	pt "github.com/TobiasGleiter/ai-agents/pkg/model_io/output/string"
)


// SimpleModel is a simple implementation of the Model interface.
type SimpleModel struct{}

func (m *SimpleModel) Process(input string) string {
	return "Model processed: " + input
}

func main() {
	// Define input
	input := "hi, bye"

	model := &SimpleModel{}
	outputParser := &pt.CommaSeparatedListOutputParser{}

	// Create a new pipe
	p := pipe.NewPipe(input, model, outputParser)
	p.Invoke()
}
