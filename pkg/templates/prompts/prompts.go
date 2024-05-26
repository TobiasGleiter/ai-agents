package prompts

import (
    "bytes"
    "text/template"
)

type PromptTemplate struct {
	template *template.Template
}

func NewPromptTemplate(templateString, templateName string) (*PromptTemplate, error) {
    tmpl, err := template.New(templateName).Parse(templateString)
    if err != nil {
        return nil, err
    }
    return &PromptTemplate{template: tmpl}, nil
}

func (pt *PromptTemplate) Format(data interface{}) (string, error) {
    var promptBuffer bytes.Buffer
    err := pt.template.Execute(&promptBuffer, data)
    if err != nil {
        return "", err
    }
    return promptBuffer.String(), nil
}