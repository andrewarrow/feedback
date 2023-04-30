package router

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
)

func (r *Router) sendZippy(doZip bool, name string, vars any, writer http.ResponseWriter, status int) {
	t := r.Template.Lookup(name)
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
