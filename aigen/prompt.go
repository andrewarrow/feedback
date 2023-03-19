package aigen

import (
	"encoding/json"
	"fmt"

	"github.com/andrewarrow/feedback/network"
)

func RunPrompt(prompt string) {

	//prompt := "User: Hi\nAI: Hello, how can I help you today?\nUser: Can you tell me more about your product?\nAI: Sure, our product is a software tool that helps businesses streamline their operations.\nUser: What are the features of your product?"

	messages := MakeMessages("Hello!")
	m := map[string]any{"model": "gpt-3.5-turbo",
		"messages": messages}
	asBytes, _ := json.Marshal(m)

	jsonString := network.DoPost("/v1/chat/completions", asBytes)
	fmt.Println(jsonString)

}

func MakeMessages(content string) []map[string]string {
	messages := []map[string]string{}
	messages = append(messages, map[string]string{"role": "user"})
	messages = append(messages, map[string]string{"content": content})
	return messages
}
