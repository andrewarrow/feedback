package router

import (
	"embed"
	"fmt"
	"net/http"
	"strings"
)

var EmbeddedAssets embed.FS

func (r *Router) HandleAsset(path string, writer http.ResponseWriter, request *http.Request) {
	contentType := "text/css"
	contentEncoding := "identity"
	if strings.HasSuffix(path, ".js") || strings.HasSuffix(path, ".json") {
		contentType = "application/javascript"
	} else if strings.HasSuffix(path, ".ico") {
		contentType = "image/x-icon"
	} else if strings.HasSuffix(path, ".gif") {
		contentType = "image/gif"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg"
	} else if strings.HasSuffix(path, ".svg") {
		contentType = "image/svg+xml"
	} else if strings.HasSuffix(path, ".ttf") {
		contentType = "font/ttf"
	} else if strings.HasSuffix(path, ".xml") {
		contentType = "text/xml"
	} else if strings.HasSuffix(path, ".wasm.gz") {
		contentEncoding = "gzip"
		contentType = "application/wasm"
	} else if strings.HasSuffix(path, ".wasm") {
		contentType = "application/wasm"
	}
	writer.Header().Set("Content-Type", contentType)
	writer.Header().Set("Connection", "keep-alive")
	writer.Header().Set("Content-Encoding", contentEncoding)
	writer.Header().Set("Cache-Control", "max-age=3600, public, must-revalidate, proxy-revalidate")
	//	matchFile := files.ReadFile(fmt.Sprintf("%s", path[1:]))

	matchFile, _ := EmbeddedAssets.ReadFile(fmt.Sprintf("%s", path[1:]))

	writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(matchFile)))
	writer.Write([]byte(matchFile))
}
