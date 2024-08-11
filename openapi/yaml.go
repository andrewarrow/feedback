package openapi

import (
	"fmt"
	"strings"
)

func MakeYaml(m map[string][]Endpoint) {

	buffer := []string{}

	for k, v := range m {
		buffer = append(buffer, "  "+k+":")
		for _, item := range v {
			buffer = append(buffer, "    "+item.LowerVerb+":")
			buffer = append(buffer, "      summary: TBD")
			if item.Method == "POST" {
				buffer = append(buffer, "      "+post)
				for _, param := range item.Params {
					buffer = append(buffer, "                "+param.Name+":")
					buffer = append(buffer, "                  type: "+param.Flavor)
				}
			}
			if item.HasId {
				buffer = append(buffer, "      parameters:")
				/*
				 parameters:
				        - name: id
				          in: path
				          required: true
				          schema:
				            type: integer
				            example: 1
				*/
			}
		}
	}

	final := yaml + "\n" + strings.Join(buffer, "\n")

	fmt.Println(final)
}

var yaml = `openapi: 3.1.0
info:
  title: Feedback API
  description: Feedback API 
  version: 1.0.0
paths:`

var post = `requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:`
