package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/util"
)

func (m *Model) CurlPutPayload() string {
	payload := map[string]any{}
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
		payload[name] = exampleVal(name, field.Flavor)
	}
	asBytes, _ := json.Marshal(payload)
	jsonString := string(asBytes)
	tokens := strings.Split(jsonString, ",")
	return "'" + strings.Join(tokens, ",\n") + "'"
}

func (m *Model) CurlPostPayload() string {
	payload := map[string]any{}
	for _, field := range m.Fields {
		if field.Required == "" {
			continue
		}
		name := field.Name
		if strings.HasSuffix(name, "_id") {
			tokens := strings.Split(name, "_")
			name = tokens[0] + "_guid"
		}
		payload[name] = exampleVal(name, field.Flavor)
	}
	asBytes, _ := json.Marshal(payload)
	jsonString := string(asBytes)
	tokens := strings.Split(jsonString, ",")
	return "'" + strings.Join(tokens, ",\n") + "'"
}

func (m *Model) CurlResponse() string {
	wrapper := map[string]any{}
	payload := m.ExampleAlFields()

	modelName := util.ToSnakeCase(m.Name)
	util.RemoveSensitiveKeys(payload)
	wrapper[modelName] = payload

	asBytes, _ := json.Marshal(wrapper)
	return util.PipeToJq(string(asBytes))
}

func (m *Model) ExampleAlFields() map[string]any {
	payload := map[string]any{}
	for _, field := range m.Fields {
		name := util.ToSnakeCase(field.Name)
		if strings.HasSuffix(name, "_id") {
			tokens := strings.Split(name, "_")
			name = tokens[0] + "_guid"
		}
		payload[name] = exampleVal(name, field.Flavor)
	}
	return payload
}

func exampleVal(name, flavor string) any {
	var val any
	if strings.HasSuffix(name, "_guid") || name == "guid" {
		val = util.PseudoUuid()
	} else if flavor == "timestamp" {
		val = time.Now().Unix()
	} else if flavor == "int" {
		val = 0
	} else {
		val = "some_string"
	}
	return val
}
