package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/util"
)

func (m *Model) CurlPutPayload() string {
	buffer := []string{}
	q := `"`
	for _, field := range m.Fields {
		if field.CommonExclude() {
			continue
		}
		if field.Required == "yes" {
			continue
		}
		name := field.Name
		if strings.HasSuffix(name, "_id") {
			continue
		}
		val := q + "some_string" + q
		if field.Flavor == "timestamp" {
			val = fmt.Sprintf("%d", time.Now().Unix())
		}

		buffer = append(buffer, fmt.Sprintf("%s%s%s: %s", q, name, q, val))
	}
	list := strings.Join(buffer, ",")
	return "'{" + list + "}'"
}

func (m *Model) CurlPostPayload() string {
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
		val := q + "some_string" + q
		if strings.HasSuffix(name, "_guid") {
			val = q + util.PseudoUuid() + q
		} else if field.Flavor == "timestamp" {
			val = fmt.Sprintf("%d", time.Now().Unix())
		}

		buffer = append(buffer, fmt.Sprintf("%s%s%s: %s", q, name, q, val))
	}
	list := strings.Join(buffer, ",")
	return "'{" + list + "}'"
}

func (m *Model) CurlResponse() string {
	wrapper := map[string]any{}
	payload := map[string]any{}

	for _, field := range m.Fields {
		payload[util.ToSnakeCase(field.Name)] = "hi"
	}

	modelName := util.ToSnakeCase(m.Name)
	wrapper[modelName] = payload

	asBytes, _ := json.Marshal(wrapper)
	return util.PipeToJq(string(asBytes))
}
