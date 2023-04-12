package util

import (
	"os/exec"
	"strings"
)

func PipeToJq(inputString string) string {

	jq := exec.Command("jq", ".")
	jq.Stdin = strings.NewReader(inputString)

	b, _ := jq.CombinedOutput()
	return string(b)

}

func GetJsonString(m map[string]any, key string) string {
	s := m[key]
	if s == nil {
		return ""
	}
	return s.(string)
}

func GetJsonMap(m map[string]any, key string) map[string]any {
	item := m[key]
	if item == nil {
		return map[string]any{}
	}
	return item.(map[string]any)
}

func GetJsonInt64(m map[string]any, key string) int64 {
	val := m[key]
	if val == nil {
		return 0
	}
	return int64(m[key].(float64))
}
