package models

import (
	"fmt"
	"strings"
)

func (m *Model) CurlCreatePayload() string {
	buffer := []string{}
	q := `"`
	for _, field := range m.Fields {
		if field.Required == "" {
			continue
		}
		name := field.Name
		if strings.HasSuffix(name, "_id") {
			tokens := strings.Split(name, "_")
			name = tokens[0] + "_guid"
		}
		val := "hi"
		buffer = append(buffer, fmt.Sprintf("%s%s%s: %s%s%s", q, name, q, q, val, q))
	}
	list := strings.Join(buffer, ",")
	return "'{" + list + "}'"
}
