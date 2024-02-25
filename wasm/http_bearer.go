package wasm

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func DoHttpBearerRead(bearer string, request *http.Request) (string, int) {
	client := &http.Client{Timeout: time.Second * 5}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer: "+bearer)
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

func DoBearerGet(bearer, urlString string) (string, int) {
	request, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return ""
	}

	jsonString, code := DoHttpBearerRead(bearer, request)
	return jsonString, code
}
