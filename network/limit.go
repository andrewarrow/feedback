package network

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func DoHttpLimitRead(client *http.Client, request *http.Request) (string, int, string) {
	const maxBodySize = 10 * 1024 * 3

	resp, err := client.Do(request)
	if err == nil {
		defer resp.Body.Close()

		limitReader := io.LimitReader(resp.Body, maxBodySize)
		body, err := io.ReadAll(limitReader)
		if err != nil {
			fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, err.Error())
			return err.Error(), 500, ""
		}

		contentLength := resp.Header.Get("Content-Length")
		return DoReadZipped(body), resp.StatusCode, contentLength
	}

	fmt.Printf("\n\nERROR: %s\n\n", err.Error())
	return err.Error(), 500, ""
}

func DoReadZipped(asBytes []byte) string {
	buf := bytes.NewBuffer(asBytes)
	gr, err := gzip.NewReader(buf)
	if err != nil {
		//fmt.Println(err)
		return ""
	}
	defer gr.Close()
	body, err := ioutil.ReadAll(gr)
	if err != nil {
		//fmt.Println(err)
		return ""
	}
	return string(body)
}

func DoHttpZRead(client *http.Client, request *http.Request, cb func(b []byte)) {
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	//contentLength := resp.Header.Get("Content-Length")
	//fmt.Println(resp.StatusCode, contentLength)
	defer resp.Body.Close()
	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		fmt.Println("Error creating Gzip reader:", err)
		return
	}
	defer reader.Close()

	chunkSize := 10 * 1024 * 3
	buffer := make([]byte, chunkSize)
	count := 0
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading from Gzip reader:", err)
			return
		}
		cb(buffer[:n])
		count++
		if count > 9 {
			break
		}
	}
}
