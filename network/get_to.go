package network

import (
	"net/http"
	"strings"
	"time"
)

func GetTo(full, bearer string) (string, int) {
	request, err := http.NewRequest("GET", full, nil)
	if err != nil {
		return "bad url", 500
	}
	SetHeaders(bearer, request)
	client := &http.Client{Timeout: time.Second * 150}

	return DoHttpRead(client, request)
}

func Get200Image(full, bearer string) bool {
	request, err := http.NewRequest("GET", full, nil)
	if err != nil {
		return false
	}
	SetHeaders(bearer, request)
	client := &http.Client{Timeout: time.Second * 150}

	resp, err := client.Do(request)
	if err != nil {
		return false
	}
	resp.Body.Close()
	contentType := strings.ToLower(resp.Header.Get("Content-Type"))
	imageName := strings.ToLower(full)
	imageInName := strings.Contains(imageName, ".jpg") || strings.Contains(imageName, ".jpeg") || strings.Contains(imageName, ".gif") || strings.Contains(imageName, ".png") || strings.Contains(imageName, ".svg") || strings.Contains(imageName, ".webp")
	if strings.Contains(contentType, "image") || imageInName {
		return resp.StatusCode == 200
	}
	return false
}
