package main

import (
	"fmt"
	pt "github.com/TobiasGleiter/ai-agents/pkg/model_io/output/string"
)

type Model struct{}

func (m *Model) Process(input string) string {
	return "Model processed: " + input
}


type OutputParser interface {
	Parse(inputs ...interface{}) ([]string, error)
}

type Pipe struct {
	Input        string
	Model        *Model
	OutputParser OutputParser
}

func (p *Pipe) Invoke() string {
	output := p.Model.Process(p.Input)
	return p.OutputParser.Parse(output)
}

func main() {
	// Define input
	input := "hi, bye"

	// Define operators
	model := &Model{}
	outputParser := &pt.CommaSeparatedListOutputParser{}

	// Define pipeline
	pipe := &Pipe{
		Input:        input,
		Model:        model,
		OutputParser: outputParser,
	}

	// Run the pipeline and print the final output
	fmt.Println(pipe.Invoke())
}
