package network

import (
	"fmt"
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
		fmt.Println(err)
		return false
	}
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("Accept-Language", "en-US,en;q=0.5")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:120.0) Gecko/20100101 Firefox/120.0")
	client := &http.Client{Timeout: time.Second * 3}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return false
	}
	resp.Body.Close()
	contentType := strings.ToLower(resp.Header.Get("Content-Type"))
	//fmt.Println("contentType", contentType)
	imageName := strings.ToLower(full)
	imageInName := strings.Contains(imageName, ".jpg") || strings.Contains(imageName, ".jpeg") || strings.Contains(imageName, ".gif") || strings.Contains(imageName, ".png") || strings.Contains(imageName, ".svg") || strings.Contains(imageName, ".webp")
	if strings.Contains(contentType, "image") || imageInName {
		return resp.StatusCode == 200
	}
	return false
}
