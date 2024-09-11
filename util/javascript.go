package util

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
)

func Minify(dir string) {
	outputFile := "assets/javascript/main.js"
	inputFiles := listOfJavascriptFiles()

	var combinedJS strings.Builder
	for _, file := range inputFiles {
		content, _ := ioutil.ReadFile(file)
		combinedJS.Write(content)
		combinedJS.WriteString("\n")
	}

	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)

	minified, _ := m.String("text/javascript", combinedJS.String())
	ioutil.WriteFile(outputFile, []byte(minified), 0644)
}

func listOfJavascriptFiles(dir string) []string {
	files, _ := os.ReadDir(dir)
	paths := []string{}
	for _, file := range files {
		name := file.Name()
		if name == "main.js" {
			continue
		}
		if strings.HasSuffix(name, ".js") == false {
			continue
		}
		paths = append(paths, dir+"/"+name)
	}
	return paths
}
