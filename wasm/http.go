package wasm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func DoHttpRead(request *http.Request) (string, int) {
	client := &http.Client{Timeout: time.Second * 5}
	request.Header.Set("Content-Type", "application/json")
	//request.Header.Set("Accept-Encoding", "application/json")
	resp, err := client.Do(request)
	if err == nil {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, err.Error())
			return err.Error(), 500
		}
		return string(body), resp.StatusCode
	}
	fmt.Printf("\n\nERROR: %s\n\n", err.Error())
	return err.Error(), 500
}

func DoGetMap(urlString string) map[string]any {
	jsonString, _ := DoGet(urlString)
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	return m
}
func DoGetItems(urlString string) []map[string]any {
	jsonString, _ := DoGet(urlString)
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	items := m["items"].([]any)
	done := []map[string]any{}
	for _, thing := range items {
		item := thing.(map[string]any)
		done = append(done, item)
	}
	return done
}

func DoGet(urlString string) (string, int) {
	request, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return "", 422
	}

	jsonString, code := DoHttpRead(request)
	return jsonString, code
}

func DoPatch(urlString string, payload any) (string, int) {
	asBytes, _ := json.Marshal(payload)
	body := bytes.NewBuffer(asBytes)
	request, err := http.NewRequest("PATCH", urlString, body)
	if err != nil {
		return err.Error(), 500
	}

	s, code := DoHttpRead(request)
	return s, code
}
func DoPut(urlString string, payload any) (string, int) {
	asBytes, _ := json.Marshal(payload)
	body := bytes.NewBuffer(asBytes)
	request, err := http.NewRequest("PUT", urlString, body)
	if err != nil {
		return "", 500
	}

	s, code := DoHttpRead(request)
	return s, code
}

func DoPost(urlString string, payload any) (string, int) {
	asBytes, _ := json.Marshal(payload)
	body := bytes.NewBuffer(asBytes)
	request, err := http.NewRequest("POST", urlString, body)
	if err != nil {
		return "", 500
	}

	s, code := DoHttpRead(request)
	return s, code
}

func DoDelete(urlString string) int {
	request, err := http.NewRequest("DELETE", urlString, nil)
	if err != nil {
		return 500
	}

	_, code := DoHttpRead(request)
	return code
}
