package network

import (
	"net/http"
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
