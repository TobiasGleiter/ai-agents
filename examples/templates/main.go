package main

import (
    "fmt"

	pt "github.com/TobiasGleiter/ai-agents/pkg/templates/prompts"
)

type Product struct {
	Product string
}

type BuildProduct struct {
	Name string
	Product string
}

func main() {
	companyNamePrompt, _ := pt.NewPromptTemplate("What is a good name for a company that makes {{.Product}}?", "company_name")
	twoVariablesPrompt, _ := pt.NewPromptTemplate("{{.Name}} want's to build {{.Product}}.", "two_variables")

	data := Product{ Product: "coloful socks" }
	companyNameFormattedPrompt, _ := companyNamePrompt.Format(data)
	fmt.Println(companyNameFormattedPrompt)

	twoVariablesData := BuildProduct{ Name: "Tobi", Product: "coloful socks"}
	twoVariablesFormattedPrompt, _ := twoVariablesPrompt.Format(twoVariablesData)
	fmt.Println(twoVariablesFormattedPrompt)
}
