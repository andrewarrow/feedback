package aigen

import (
	"encoding/json"
	"fmt"

	"github.com/andrewarrow/feedback/network"
)

func RunPrompt(prompt string) {

	//prompt := "User: Hi\nAI: Hello, how can I help you today?\nUser: Can you tell me more about your product?\nAI: Sure, our product is a software tool that helps businesses streamline their operations.\nUser: What are the features of your product?"

	messages := MakeMessages(`you are an issue tracking system that uses terms like Project, Epic, Story, Task, and Bug. When I ask you to do something to the system in english, respond with json like {"steps": ["insert_issue": ["title", "bug"], "assign_issue": "username", "update_issue": "title=x"]}. Be sure to include N number of steps. If the user is asking to do multiple things, don't assume they can be done all in one step. For example the user asks to assign an issue to a user, make a seperate step to create that user. Don't include any sql in your response, just nice json that explains what sql will need to be crafted.  make a new issue with title bug with encoding video and assign it to mark then add it to the current epic.`)
	m := map[string]any{"model": "gpt-3.5-turbo",
		"messages": messages}
	asBytes, _ := json.Marshal(m)

	jsonString := network.DoPost("/v1/chat/completions", asBytes)
	fmt.Println(jsonString)

}

func MakeMessages(content string) []map[string]string {
	messages := []map[string]string{}
	message := map[string]string{"role": "user", "content": content}
	messages = append(messages, message)
	return messages
}
