package router

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func (c *Context) SendContentAsJson(thing any, status int) {
	asBytes, _ := json.Marshal(thing)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Cache-Control", "none")
	c.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(asBytes)))
	c.Writer.WriteHeader(status)
	c.Writer.Write(asBytes)
}

func (c *Context) JsonInfo(message string) map[string]any {
	m := map[string]any{"info": message}
	return m
}

func (c *Context) ValidateJsonForModel(modelString string) string {
	params := c.ReadBodyIntoJson()
	model := c.FindModel(modelString)

	for _, field := range model.RequiredFields() {
		if params[field.Name] == nil {
			return "missing " + field.Name
		}
	}

	for _, field := range model.Fields {
		if field.Regex == "" {
			continue
		}
		if params[field.Name] == nil {
			continue
		}

		val := params[field.Name].(string)

		re := regexp.MustCompile(field.Regex)
		if !re.MatchString(val) {
			return "wrong format " + field.Name
		}
	}

	return ""
}
