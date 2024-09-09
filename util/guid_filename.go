package util

import "strings"

func GuidFilename(name, guid string) string {
	if !strings.Contains(name, ".") {
		name = name + ".bin"
	}
	tokens := strings.Split(name, ".")
	ext := tokens[len(tokens)-1]
	filename := guid + "." + ext
	return filename
}
