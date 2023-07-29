package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (r *Router) HandleBucketAsset(path string, writer http.ResponseWriter, request *http.Request) {
	contentType := "text/css"
	if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else if strings.HasSuffix(path, ".ico") {
		contentType = "image/x-icon"
	} else if strings.HasSuffix(path, ".gif") {
		contentType = "image/gif"
	} else if strings.HasSuffix(path, ".pdf") {
		contentType = "application/pdf"
	}
	writer.Header().Set("Content-Type", contentType)
	writer.Header().Set("Connection", "keep-alive")
	writer.Header().Set("Cache-Control", "max-age=3600, public, must-revalidate, proxy-revalidate")

	asBytes, _ := ioutil.ReadFile(r.BucketPath + path)
	writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(asBytes)))
	writer.Write(asBytes)
}
