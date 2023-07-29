package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (r *Router) HandleBucketAsset(path string, writer http.ResponseWriter, request *http.Request) {
	contentType := "application/octet-stream"
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
	path = path[7:]
	asBytes, _ := ioutil.ReadFile(r.BucketPath + path)
	writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(asBytes)))
	writer.Write(asBytes)
}
