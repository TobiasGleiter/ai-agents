package main

import (
	"fmt"

	pt "github.com/TobiasGleiter/ai-agents/pkg/model_io/output"
)

func main() {
	outputParser := pt.CommaSeparatedListOutputParser{}

	parsedOutput := outputParser.Parse("hi, bye")

	fmt.Println(parsedOutput)
	fmt.Println(parsedOutput[0])
}