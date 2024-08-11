package openapi

import (
	"io/ioutil"
	"sort"
	"strings"
)

func comparePaths(path1, path2 string) bool {
	tokens1 := strings.Split(path1, "/")
	tokens2 := strings.Split(path2, "/")

	for i := 0; i < len(tokens1) && i < len(tokens2); i++ {
		if tokens1[i] != tokens2[i] {
			return tokens1[i] > tokens2[i]
		}
	}

	return len(tokens1) < len(tokens2)
}

func (oa *OpenAPI) WriteYaml() {
	buffer := []string{}

	items := []string{}
	for k, _ := range oa.Endpoints {
		items = append(items, k)
	}
	sort.Slice(items, func(i, j int) bool {
		return comparePaths(items[i], items[j])
	})
	for i := len(items) - 1; i >= 0; i-- {
		k := items[i]
		v := oa.Endpoints[k]
		buffer = append(buffer, "  /"+k+":")
		for _, item := range v {
			buffer = append(buffer, "    "+item.LowerVerb+":")
			buffer = append(buffer, "      summary: ...")
			if item.Method == "POST" {
				buffer = append(buffer, "      "+post)
				for _, param := range oa.ParamsByFunc[item.CallingFunc] {
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
