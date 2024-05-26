package pipe

import "fmt"

// OutputParser is the interface that different parsers should implement.
type OutputParser interface {
	Parse(output string) interface{}
}

// Model represents the machine learning model.
type Model interface {
	Process(input string) string
}

// Pipe represents the pipeline with input, model, and output parser.
type Pipe struct {
	Input        string
	Model        Model
	OutputParser OutputParser
}

func NewPipe(input string, model Model, outputParser OutputParser) *Pipe {
	return &Pipe{
		Input:        input,
		Model:        model,
		OutputParser: outputParser,
	}
}

func (p *Pipe) Invoke() {
	output := p.Model.Process(p.Input)
	parsedOutput := p.OutputParser.Parse(output)
	
	switch po := parsedOutput.(type) {
	case []string:
		fmt.Println("Parsed output as slice of strings:", po)
	case string:
		fmt.Println("Parsed output as string:", po)
	default:
		fmt.Println("Unknown output type")
	}
}
