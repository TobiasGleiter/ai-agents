package messages

import (
    "bytes"
	"text/template"
)

type ChatMessage struct {
    Role    string
    Content string
}

type ChatPromptTemplate struct {
    Messages []ChatMessage
}

func NewChatPromptTemplate(messages []ChatMessage) (*ChatPromptTemplate, error) {
    return &ChatPromptTemplate{Messages: messages,}, nil
}

func (cpt *ChatPromptTemplate) FormatMessages(data map[string]interface{}) ([]ChatMessage, error) {
    var formattedMessages []ChatMessage

    for _, templat := range cpt.Messages {
        tmpl, err := template.New("prompt").Parse(templat.Content)
        if err != nil {
            return nil, err
        }

        var buffer bytes.Buffer
        err = tmpl.Execute(&buffer, data)
        if err != nil {
            return nil, err
        }

        formattedMessages = append(formattedMessages, ChatMessage{
            Role:    templat.Role,
            Content: buffer.String(),
        })
    }

    return formattedMessages, nil
}