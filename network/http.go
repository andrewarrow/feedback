package network

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

var BaseUrl = "https://api.openai.com"

func DoGet(bearer, route string) (string, int) {
	urlString := fmt.Sprintf("%s%s", BaseUrl, route)
	request, _ := http.NewRequest("GET", urlString, nil)
	SetHeaders(bearer, request)
	client := &http.Client{Timeout: time.Second * 5}
	jsonString, code := DoHttpRead(client, request)
	return jsonString, code
}

func DoHttpRead(client *http.Client, request *http.Request) (string, int) {
	resp, err := client.Do(request)
	if err == nil {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		//var buff bytes.Buffer
		//io.Copy(&buff, resp.Body)
		//body := buff.Bytes()
		if err != nil {
			fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, err.Error())
			return err.Error(), 500
		}
		return string(body), resp.StatusCode
	}
	fmt.Printf("\n\nERROR: %s\n\n", err.Error())
	return err.Error(), 500
}

func DoPost(bearer, route string, payload []byte) (string, int) {
	body := bytes.NewBuffer(payload)
	urlString := fmt.Sprintf("%s%s", BaseUrl, route)
	request, _ := http.NewRequest("POST", urlString, body)
	SetHeaders(bearer, request)
	client := &http.Client{Timeout: time.Second * 50}

	jsonString, code := DoHttpRead(client, request)
	return jsonString, code
}

func DoPut(bearer, route string, payload []byte) (string, int) {
	body := bytes.NewBuffer(payload)
	urlString := fmt.Sprintf("%s%s", BaseUrl, route)
	request, _ := http.NewRequest("PUT", urlString, body)
	SetHeaders(bearer, request)
	client := &http.Client{Timeout: time.Second * 50}

	jsonString, code := DoHttpRead(client, request)
	return jsonString, code
}

func DoDelete(bearer, route string) (string, int) {
	urlString := fmt.Sprintf("%s%s", BaseUrl, route)
	request, _ := http.NewRequest("DELETE", urlString, nil)
	SetHeaders(bearer, request)
	client := &http.Client{Timeout: time.Second * 50}

	jsonString, code := DoHttpRead(client, request)
	return jsonString, code
}

func DoMultiPartPost(bearer, route, name string, payload []byte) (string, int) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	//writer.WriteField("name", "John Doe")

	fileWriter, _ := writer.CreateFormFile(name, "multipart")
	theData := bytes.NewBuffer(payload)
	io.Copy(fileWriter, theData)

	writer.Close()

	urlString := fmt.Sprintf("%s%s", BaseUrl, route)
	request, _ := http.NewRequest("POST", urlString, body)

	SetHeaders(bearer, request)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{Timeout: time.Second * 50}

	jsonString, code := DoHttpRead(client, request)
	return jsonString, code
}

func SetHeaders(bearer string, request *http.Request) {
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearer))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
}
