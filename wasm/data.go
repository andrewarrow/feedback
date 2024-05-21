package wasm

import (
	"encoding/json"
	"strings"
)

func (g *Global) LoadData(route, guid string) []any {
	tokens := strings.Split(guid, "-")
	tokens = tokens[1:]
	id := strings.Join(tokens, "-")
	jsonString := DoGet(route + "&guid=" + id)
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	items, _ := m["items"].([]any)
	return items
}
