package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/files"
)

func (r *Router) HandleAsset(path string, writer http.ResponseWriter) {
	contentType := "text/css"
	if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else if strings.HasSuffix(path, ".gif") {
		contentType = "image/gif"
	}
	writer.Header().Set("Content-Type", contentType)
	matchFile := files.ReadFile(fmt.Sprintf("%s", path[1:]))
	fmt.Fprintf(writer, matchFile)
}
