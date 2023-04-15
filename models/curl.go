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
		payload[name] = exampleVal(name, field)
	}
	asBytes, _ := json.Marshal(payload)
	jsonString := string(asBytes)
	return "'" + util.PipeToJq(jsonString) + "'"
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
		payload[name] = exampleVal(name, field)
	}
	asBytes, _ := json.Marshal(payload)
	jsonString := string(asBytes)
	return "'" + util.PipeToJq(jsonString) + "'"
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
		payload[name] = exampleVal(name, field)
	}
	return payload
}

func exampleVal(name string, field *Field) any {
	var val any
	if strings.HasSuffix(name, "_guid") || name == "guid" {
		val = util.PseudoUuid()
	} else if field.Flavor == "timestamp" {
		val = time.Now().Unix()
	} else if field.Flavor == "int" {
		val = 0
	} else if len(field.JsonNames) > 0 {
		val = makeJsonFromNames(field)
	} else {
		val = "some_string"
	}
	return val
}

func makeJsonFromNames(field *Field) map[string]any {
	payload := map[string]any{}
	payload["foo"] = 123
	return payload
}
