package main

import (
    "fmt"

	pt "github.com/TobiasGleiter/ai-agents/pkg/model_io/templates/prompt"
)

type Product struct {
	Product string
}

type BuildProduct struct {
	Name string
	Company string
}

func main() {
	companyNamePrompt, _ := pt.NewPromptTemplate("What is a good name for a company that makes {{.Product}}?", "company_name")
	twoVariablesPrompt, _ := pt.NewPromptTemplate("{{.Name}} want's to build {{.Company}}.", "two_variables")

	data := Product{ Product: "coloful socks" }
	companyNameFormattedPrompt, _ := companyNamePrompt.Format(data)
	fmt.Println(companyNameFormattedPrompt)

	buildProduct := BuildProduct{ Name: "Tobi", Company: "coloful socks"}
	twoVariablesFormattedPrompt, _ := twoVariablesPrompt.Format(buildProduct)
	fmt.Println(twoVariablesFormattedPrompt)
}
