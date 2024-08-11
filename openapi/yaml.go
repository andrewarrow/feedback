package openapi

import (
	"io/ioutil"
	"sort"
	"strings"
)

func (oa *OpenAPI) WriteYaml() {
	buffer := []string{}

	items := []string{}
	for k, _ := range oa.Endpoints {
		items = append(items, k)
	}
	sort.Strings(items)
	for _, k := range items {
		v := oa.Endpoints[k]
		buffer = append(buffer, "  /"+k+":")
		for _, item := range v {
			buffer = append(buffer, "    "+item.LowerVerb+":")
			buffer = append(buffer, "      summary: ...")
			if item.Method == "POST" {
				buffer = append(buffer, "      "+post)
				for _, param := range item.Params {
					buffer = append(buffer, "                "+param.Name+":")
					buffer = append(buffer, "                  type: "+param.Flavor)
				}
			}
			if item.HasId {
				buffer = append(buffer, "      parameters:")
				buffer = append(buffer, "        - name: id")
				buffer = append(buffer, "          in: path")
				buffer = append(buffer, "          required: true")
				buffer = append(buffer, "          schema:")
				buffer = append(buffer, "            type: string")
			}
			buffer = append(buffer, "      responses:")
			buffer = append(buffer, "        '200':")
			buffer = append(buffer, "          description: ok")
			buffer = append(buffer, "          content:")
			buffer = append(buffer, "            application/json:")
			buffer = append(buffer, "              schema:")
			buffer = append(buffer, "                type: object")
			buffer = append(buffer, "                properties:")
			buffer = append(buffer, "                  msg:")
			buffer = append(buffer, "                    type: string")
		}
	}
	final := yaml + "\n" + strings.Join(buffer, "\n")

	ioutil.WriteFile("openapi/openapi.yaml", []byte(final), 0644)
}

var yaml = `openapi: 3.0.3
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
