package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/util"
	"github.com/brianvoe/gofakeit/v6"
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
	return "'" + jsonString + "'"
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
	return "'" + jsonString + "'"
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
	} else if field.Flavor == "url" {
		val = gofakeit.URL()
	} else if field.Flavor == "bool" {
		val = true
	} else if field.Flavor == "list" {
		val = []string{"item1", "item2"}
	} else if field.Flavor == "json" {
		val = makeJsonFromNames(field)
	} else {
		val = "some_string"
	}
	return val
}

func makeJsonFromNames(field *Field) map[string]any {
	payload := map[string]any{}
	var holder *Field
	for i, name := range field.JsonNames {
		f := Field{}
		f.Flavor = field.JsonTypes[i]
		holder = &f
		payload[name] = exampleVal(name, holder)
	}
	return payload
}
