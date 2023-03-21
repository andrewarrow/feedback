package util

import (
	"net/http"
	"strings"
)

func GetHeader(field string, request *http.Request) string {
	//fmt.Printf("%+v", request.Header)
	val := request.Header[field]
	if len(val) == 0 {
		return ""
	}
	return strings.Join(val, ",")
}
