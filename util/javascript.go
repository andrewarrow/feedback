package util

import (
	"os"
	"strings"
)

func Minify(dir string) {
}
func ListOfJavascriptFiles(dir string) []string {
	files, _ := os.ReadDir(dir)
	paths := []string{}
	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".js") == false {
			continue
		}
		paths = append(paths, dir+"/"+name)
	}
	return paths
}
