package main

import (
	"fmt"

	pt "github.com/TobiasGleiter/ai-agents/pkg/model_io/output/string"
	jpt "github.com/TobiasGleiter/ai-agents/pkg/model_io/output/json"
)

type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func main() {
	commaSeparatedListOutputParser := &pt.CommaSeparatedListOutputParser{}
	parsedOutput, _ := commaSeparatedListOutputParser.Parse("hi, bye")
	fmt.Println(parsedOutput)
	fmt.Println(parsedOutput[0])

	
	var joke Joke
	jsonOutputParser := &jpt.JsonOutputParser{}
	err := jsonOutputParser.Parse(`{"setup": "Why don't scientists trust atoms?", "punchline": "Because they make up everything!"}`, &joke)
	if err != nil {
		panic(err)
	}
	fmt.Println(joke.Setup)
	fmt.Println(joke.Punchline)
}