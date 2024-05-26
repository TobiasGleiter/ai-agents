package main

import (
    "fmt"

	pt "github.com/TobiasGleiter/ai-agents/pkg/templates/messages"
)


func main() {
	chatPrompt, _ := pt.NewChatPromptTemplate([]pt.ChatMessageTemplate{
        {Role: "system", Content: "You are a helpful assistant that translates {{.InputLanguage}} to {{.OutputLanguage}}."},
        {Role: "user", Content: "{{.Text}}"},
    })

	data := map[string]interface{}{
        "InputLanguage":  "English",
        "OutputLanguage": "French",
        "Text":           "I love programming.",
    }

	formattedMessages, err := chatPrompt.FormatMessages(data)
    if err != nil {
        panic(err)
    }

	for _, message := range formattedMessages {
        fmt.Printf("[%s] %s\n", message.Role, message.Content)
    }
}
