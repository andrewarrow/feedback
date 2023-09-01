package location

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ReadInZips(dirPath string) {
	files, _ := ioutil.ReadDir(dirPath)
	for _, file := range files {
		filename := dirPath + "/" + file.Name()
		ReadInZipsState(filename)
	}
}

func processLine(line string) {
	line = strings.TrimSpace(line)
	if line == "" {
		fmt.Println("no zip")
		return
	}
	var m map[string]any
	json.Unmarshal([]byte(line), &m)
	properties := m["properties"].(map[string]any)
	zip, ok := properties["postcode"].(string)
	if !ok {
		fmt.Println("no zip", line)
		return
	}
	geo := m["geometry"].(map[string]any)
	latlong := geo["coordinates"].([]any)
	if len(latlong) == 2 && len(zip) == 5 {
		fmt.Println(zip, latlong)
	} else {
		fmt.Println("no zip", line)
	}
}

func handleFileInBatches(filename string) {
	file, _ := os.Open(filename)
	buffer := make([]byte, 1)
	line := []string{}
	for {
		n, err := file.Read(buffer)

		if err == io.EOF || n == 0 {
			break
		}

		s := string(buffer)
		if s == "\n" {
			theLine := strings.Join(line, "")
			processLine(theLine)
			line = []string{}
		}
		line = append(line, s)
	}
	file.Close()
}

func ReadInZipsState(dirPath string) {
	// from https://openaddresses.io/
	// https://batch.openaddresses.io/data
	// alameda-addresses-county.geojson
	// {"type":"Feature","properties":{"hash":"9967fbc6d2dd1931","number":"41829","street":"OSGOOD RD","unit":"","city":"FREMONT","district":"","region":"","postcode":"94539","id":"525 034200500"},"geometry":{"type":"Point","coordinates":[-121.952505,37.5293622]}}
	files, _ := ioutil.ReadDir(dirPath)
	for _, file := range files {
		filename := dirPath + "/" + file.Name()
		if strings.HasSuffix(filename, ".meta") {
			continue
		}
		handleFileInBatches(filename)
	}
}
