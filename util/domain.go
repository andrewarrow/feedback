package util

import "strings"

func ExtractDomain(url string) string {
	tokens := strings.Split(url, "/")
	if len(tokens) > 2 {
		tokens = strings.Split(tokens[2], ".")
		if len(tokens) == 3 {
			tokens = tokens[1:]
		}
		return strings.Join(tokens, ".")
	}
	return ""
}
