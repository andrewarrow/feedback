package router

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/markup"
)

var UseLiveTemplates = false

func (r *Router) getLiveOrCachedTemplate(name string) *template.Template {
	var t *template.Template
	if UseLiveTemplates {
		if strings.HasSuffix(name, ".mu") {
			send := map[string]any{}
			rendered := markup.ToHTML(send, "markup/"+name)
			fmt.Println(rendered)
			t, _ = template.New("markup").Parse(rendered)
		} else {
			t, _ = template.ParseFiles("views/" + name)
		}
	} else {
		t = r.Template.Lookup(name)
	}
	return t
}

func (r *Router) sendZippy(doZip bool, name string, vars any, writer http.ResponseWriter, status int) {
	t := r.getLiveOrCachedTemplate(name)
	content := new(bytes.Buffer)
	t.Execute(content, vars)
	cb := content.Bytes()

	if doZip {
		writer.Header().Set("Content-Encoding", "gzip")

		var compressedData bytes.Buffer
		gzipWriter := gzip.NewWriter(&compressedData)
		gzipWriter.Write(cb)
		gzipWriter.Close()

		cb = compressedData.Bytes()
	}
	writer.Header().Set("Content-Type", "text/html")
	writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(cb)))
	writer.WriteHeader(status)
	writer.Write(cb)
}

func doZippyJson(doZip bool, asBytes []byte, status int, writer http.ResponseWriter) {
	if doZip {
		writer.Header().Set("Content-Encoding", "gzip")

		var compressedData bytes.Buffer
		gzipWriter := gzip.NewWriter(&compressedData)
		gzipWriter.Write(asBytes)
		gzipWriter.Close()

		asBytes = compressedData.Bytes()
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Cache-Control", "none")
	writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(asBytes)))
	writer.WriteHeader(status)
	writer.Write(asBytes)
}
