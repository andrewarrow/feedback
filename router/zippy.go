package router

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/andrewarrow/feedback/markup"
)

var UseLiveTemplates = os.Getenv("USE_LIVE_TEMPLATES") == "true"

func (r *Router) GetLiveOrCachedTemplate(name string) *template.Template {
	var t *template.Template
	if UseLiveTemplates {

		list, _ := ioutil.ReadDir("markup")
		for _, file := range list {
			name := file.Name()
			//fmt.Println("*", name)
			tokens := strings.Split(name, ".")
			send := map[string]any{}
			rendered := markup.ToHTML(send, "markup/"+name)
			//fmt.Println(rendered)
			ioutil.WriteFile("views/"+tokens[0]+".html", []byte(rendered), 0644)
		}
		live := LoadLiveTemplates()
		t = live.Lookup(name)
	} else {
		t = r.Template.Lookup(name)
	}
	return t
}

func (r *Router) sendZippy(doZip bool, name string, vars any, writer http.ResponseWriter, status int) {
	t := r.GetLiveOrCachedTemplate(name)
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
