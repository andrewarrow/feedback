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
	request, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return "bad url", 500
	}
	SetHeaders(bearer, request)
	client := &http.Client{Timeout: time.Second * 5}
	return DoHttpRead(client, request)
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
	request, err := http.NewRequest("POST", urlString, body)
	if err != nil {
		return "bad url", 500
	}
	SetHeaders(bearer, request)
	client := &http.Client{Timeout: time.Second * 50}

	return DoHttpRead(client, request)
}

func DoPut(bearer, route string, payload []byte) (string, int) {
	body := bytes.NewBuffer(payload)
	urlString := fmt.Sprintf("%s%s", BaseUrl, route)
	request, err := http.NewRequest("PUT", urlString, body)
	if err != nil {
		return "bad url", 500
	}
	SetHeaders(bearer, request)
	client := &http.Client{Timeout: time.Second * 50}

	return DoHttpRead(client, request)
}

func DoDelete(bearer, route string) (string, int) {
	urlString := fmt.Sprintf("%s%s", BaseUrl, route)
	request, err := http.NewRequest("DELETE", urlString, nil)
	if err != nil {
		return "bad url", 500
	}
	SetHeaders(bearer, request)
	client := &http.Client{Timeout: time.Second * 50}

	return DoHttpRead(client, request)
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
	request, err := http.NewRequest("POST", urlString, body)
	if err != nil {
		return "bad url", 500
	}

	SetHeaders(bearer, request)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{Timeout: time.Second * 50}

	return DoHttpRead(client, request)
}

func SetHeaders(bearer string, request *http.Request) {
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearer))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
}
