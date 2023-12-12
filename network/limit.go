package network

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func DoHttpLimitRead(client *http.Client, request *http.Request) (string, int) {
	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("\n\nERROR: %s\n\n", err.Error())
		return err.Error(), 500
	}
	defer resp.Body.Close()

	ce := resp.Header.Get("Content-Encoding")
	var body []byte
	const desiredSize = 100 * 1024 // 100KB
	byteValues := make([]byte, desiredSize)

	totalRead := 0
	for totalRead < desiredSize {
		n, err := resp.Body.Read(byteValues[totalRead:])
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("\n\nERROR: %d %s\n\n", resp.StatusCode, err.Error())
			return err.Error(), 500
		}
		totalRead += n
	}

	if ce == "gzip" {
		buf := bytes.NewBuffer(byteValues[:totalRead])
		gr, err := gzip.NewReader(buf)
		if err != nil {
			fmt.Printf("\n\nERROR creating gzip reader: %s\n\n", err.Error())
			return err.Error(), 500
		}
		defer gr.Close()

		body, err = ioutil.ReadAll(gr)
		if err != nil {
			fmt.Printf("\n\nERROR reading uncompressed data: %s\n\n", err.Error())
			return err.Error(), 500
		}
	} else {
		body = byteValues[:totalRead]
	}

	return string(body), resp.StatusCode
}
