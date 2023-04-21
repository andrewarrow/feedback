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

func (m *Model) CurlSingleResponseNoWrapper() string {
	asBytes, _ := json.Marshal(m.CurlResponse())
	return util.PipeToJq(string(asBytes))
}

func (m *Model) CurlSingleResponse() string {
	wrapper := map[string]any{}
	modelName := util.ToSnakeCase(m.Name)
	wrapper[modelName] = m.CurlResponse()
	asBytes, _ := json.Marshal(wrapper)
	return util.PipeToJq(string(asBytes))
}

func (m *Model) CurlListResponse() string {
	wrapper := []any{}
	wrapper = append(wrapper, m.CurlResponse())
	asBytes, _ := json.Marshal(wrapper)
	return util.PipeToJq(string(asBytes))
}

func (m *Model) CurlResponse() any {
	payload := m.ExampleAlFields()
	util.RemoveSensitiveKeys(payload)
	return payload
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
	} else if field.Flavor == "timestamp_human" {
		val = time.Now().Format(HUMAN)
	} else if field.Flavor == "int" {
		val = 0
	} else if field.Flavor == "url" {
		val = gofakeit.URL()
	} else if field.Flavor == "last4" {
		val = "1111"
	} else if field.Flavor == "brand" {
		val = "visa"
	} else if field.Flavor == "credit_card" {
		val = "4111111111111111"
	} else if field.Flavor == "year" {
		val = "2024"
	} else if field.Flavor == "month" {
		val = "10"
	} else if field.Flavor == "cvc" {
		val = "123"
	} else if field.Flavor == "bool" {
		val = true
	} else if field.Flavor == "list" {
		val = []string{"item1", "item2"}
	} else if field.Flavor == "json" {
		val = makeJsonFromNames(field)
	} else if field.Flavor == "json_list" {
		val = makeJsonListFromNames(field)
	} else {
		val = field.Flavor
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

func makeJsonListFromNames(field *Field) []map[string]any {
	payload := makeJsonFromNames(field)
	return []map[string]any{payload}
}
