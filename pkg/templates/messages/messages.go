package messages

import (
    "bytes"
	"text/template"
)

type ChatMessageTemplate struct {
    Role    string
    Content string
}

type ChatPromptTemplate struct {
    MessageTemplates []ChatMessageTemplate
}

func NewChatPromptTemplate(templates []ChatMessageTemplate) (*ChatPromptTemplate, error) {
    return &ChatPromptTemplate{MessageTemplates: templates,}, nil
}

func (cpt *ChatPromptTemplate) FormatMessages(data map[string]interface{}) ([]ChatMessageTemplate, error) {
    var formattedMessages []ChatMessageTemplate

    for _, templat := range cpt.MessageTemplates {
        tmpl, err := template.New("prompt").Parse(templat.Content)
        if err != nil {
            return nil, err
        }

        var buffer bytes.Buffer
        err = tmpl.Execute(&buffer, data)
        if err != nil {
            return nil, err
        }

        formattedMessages = append(formattedMessages, ChatMessageTemplate{
            Role:    templat.Role,
            Content: buffer.String(),
        })
    }

    return formattedMessages, nil
}