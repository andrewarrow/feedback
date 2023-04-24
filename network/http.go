package network

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/andrewarrow/feedback/router"
)

var BaseUrl = "https://api.openai.com"

func DoGet(route string) string {
	agent := "agent"

	urlString := fmt.Sprintf("%s%s", BaseUrl, route)
	request, _ := http.NewRequest("GET", urlString, nil)
	request.Header.Set("User-Agent", agent)
	request.Header.Set("Authorization", "Bearer "+os.Getenv("OPEN_AI"))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 5}
	return DoHttpRead(route, client, request)
}

func DoHttpRead(route string, client *http.Client, request *http.Request) string {
	resp, err := client.Do(request)
	if err == nil {
		defer resp.Body.Close()
		//body, err := ioutil.ReadAll(resp.Body)
		var buff bytes.Buffer
		io.Copy(&buff, resp.Body)
		body := buff.Bytes()
		if err != nil {
			fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, err.Error())
			return ""
		}
		if resp.StatusCode == 200 || resp.StatusCode == 201 || resp.StatusCode == 204 {
			return string(body)
		} else {
			text := string(body)
			fmt.Println(text)
			return ""
		}
	}
	fmt.Printf("\n\nERROR: %s\n\n", err.Error())
	return ""
}

func DoPost(route string, payload []byte) string {
	body := bytes.NewBuffer(payload)
	urlString := fmt.Sprintf("%s%s", BaseUrl, route)
	request, _ := http.NewRequest("POST", urlString, body)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+os.Getenv("OPEN_AI"))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 50}

	return DoHttpRead(route, client, request)
}

func DoMultiPartPost(c *router.Context, route, name string, payload []byte) string {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	//writer.WriteField("name", "John Doe")

	fileWriter, _ := writer.CreateFormFile(name, "multipart")
	theData := bytes.NewBuffer(payload)
	io.Copy(fileWriter, theData)

	writer.Close()

	urlString := fmt.Sprintf("%s%s", BaseUrl, route)
	request, _ := http.NewRequest("POST", urlString, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	userCookie := router.GetCookie(c, "user")
	SetHeaders(userCookie, request)
	client := &http.Client{Timeout: time.Second * 50}

	return DoHttpRead(route, client, request)
}

func SetHeaders(bearer string, request *http.Request) {
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearer))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
}
