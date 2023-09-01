package location

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadInZips(dirPath string) {
	// alameda-addresses-county.geojson
	// {"type":"Feature","properties":{"hash":"9967fbc6d2dd1931","number":"41829","street":"OSGOOD RD","unit":"","city":"FREMONT","district":"","region":"","postcode":"94539","id":"525 034200500"},"geometry":{"type":"Point","coordinates":[-121.952505,37.5293622]}}
	files, _ := ioutil.ReadDir(dirPath)
	for _, file := range files {
		filename := dirPath + "/" + file.Name()
		if strings.HasSuffix(filename, ".meta") {
			continue
		}
		//fmt.Println(filename)
		b, _ := ioutil.ReadFile(filename)
		asString := string(b)
		//fmt.Println(len(b))
		lines := strings.Split(asString, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			var m map[string]any
			json.Unmarshal([]byte(line), &m)
			properties := m["properties"].(map[string]any)
			zip, ok := properties["postcode"].(string)
			if !ok {
				continue
			}
			geo := m["geometry"].(map[string]any)
			latlong := geo["coordinates"].([]any)
			if len(latlong) == 2 && len(zip) == 5 {
				fmt.Println(zip, latlong)
			}
		}
	}
}
