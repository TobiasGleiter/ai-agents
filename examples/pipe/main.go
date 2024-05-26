package main

import (
	"fmt"

	"github.com/TobiasGleiter/ai-agents/pkg/pipe"
	pt "github.com/TobiasGleiter/ai-agents/pkg/model_io/output/string"
	"github.com/TobiasGleiter/ai-agents/pkg/core/model"
)



func main() {
	input := "hi, bye"
	model := &model.SimpleModel{}
	outputParser := &pt.CommaSeparatedListOutputParser{}

	inputModelOutputPipe := pipe.NewPipe(input, model, outputParser)
	generatedOutput := inputModelOutputPipe.Invoke()
	fmt.Println(generatedOutput)
}
