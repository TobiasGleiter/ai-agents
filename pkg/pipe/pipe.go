package pipe

type OutputParser interface {
	Parse(output string) interface{}
}

type Model interface {
	Process(input string) string
}

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

func (p *Pipe) Invoke() interface{} {
	output := p.Model.Process(p.Input)
	parsedOutput := p.OutputParser.Parse(output)
	return parsedOutput
}


