package main

import (
    "fmt"

	pt "github.com/TobiasGleiter/ai-agents/pkg/prompts"
)


func main() {
	prompt, _ := pt.NewPromptTemplate("What is a good name for a company that makes {{.Product}}?", "company_name")

	data := struct {
		Product string
	}{
		Product: "coloful socks",
	}

	formattedPrompt, _ := prompt.Format(data)

	fmt.Println(formattedPrompt)
}
