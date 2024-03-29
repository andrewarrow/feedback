package location

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/andrewarrow/feedback/router"
)

func ReadInZips(c *router.Context, dirPath string) {

	files, _ := ioutil.ReadDir(dirPath)
	for _, file := range files {
		filename := dirPath + "/" + file.Name()
		ReadInZipsState(c, filename)
	}
}

var gcount = 0

func processLine(c *router.Context, line string) {
	gcount++
	if gcount%1000 == 0 {
		fmt.Println(gcount, line)
	}

	line = strings.TrimSpace(line)
	if line == "" {
		//fmt.Println("no zip")
		return
	}
	var m map[string]any
	json.Unmarshal([]byte(line), &m)
	properties := m["properties"].(map[string]any)
	zip, ok := properties["postcode"].(string)
	if !ok {
		//fmt.Println("no zip", line)
		return
	}
	geo, ok := m["geometry"].(map[string]any)
	if !ok {
		//fmt.Println("no zip", line)
		return
	}
	latlong := geo["coordinates"].([]any)
	if len(latlong) == 2 && len(zip) == 5 {
		//fmt.Println(zip, latlong)
		s1 := `INSERT INTO zip_locations (zip, location) VALUES (%d, %s)`
		s2 := `ST_SetSRID(ST_MakePoint(%f, %f), 4326)`
		loc := fmt.Sprintf(s2, latlong[0], latlong[1])
		zipInt, _ := strconv.Atoi(zip)
		sql := fmt.Sprintf(s1, zipInt, loc)
		//fmt.Println(sql)
		c.FreeFormUpdate(sql)
	} else {
		//fmt.Println("no zip", line)
	}
}

func handleFileInBatches(c *router.Context, filename string) {
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
			processLine(c, theLine)
			line = []string{}
		}
		line = append(line, s)
	}
	file.Close()
}

func ReadInZipsState(c *router.Context, dirPath string) {
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
		handleFileInBatches(c, filename)
	}
}
