package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/files"
)

func (r *Router) HandleAsset(path string, writer http.ResponseWriter, request *http.Request) {
	contentType := "text/css"
	if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else if strings.HasSuffix(path, ".ico") {
		contentType = "image/x-icon"
	} else if strings.HasSuffix(path, ".gif") {
		contentType = "image/gif"
	}
	writer.Header().Set("Content-Type", contentType)
	writer.Header().Set("Connection", "keep-alive")
	writer.Header().Set("Cache-Control", "max-age=3600, public, must-revalidate, proxy-revalidate")
	matchFile := files.ReadFile(fmt.Sprintf("%s", path[1:]))
	writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(matchFile)))
	writer.Write([]byte(matchFile))
}
