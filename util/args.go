package util

import (
	"os"
	"strconv"
	"strings"
)

func GetArg(index int) string {
	if len(os.Args) > index {
		return os.Args[index]
	}
	return ""
}

func Atoi(num string, ifzero int) int {
	thing, _ := strconv.Atoi(num)
	if thing == 0 {
		return ifzero
	}
	return thing
}

func ArgsToMap() map[string]string {
	m := map[string]string{}
	if len(os.Args) == 1 {
		return m
	}

	for _, a := range os.Args[1:] {
		if strings.HasPrefix(a, "--") {
			tokens := strings.Split(a, "=")
			key := strings.Split(tokens[0], "--")
			if len(tokens) == 2 {
				m[key[1]] = tokens[1]
			} else {
				m[key[1]] = "true"
			}
		} else if strings.Contains(a, "=") {
			tokens := strings.Split(a, "=")
			if len(tokens) == 2 {
				m[tokens[0]] = tokens[1]
			}
		}
	}
	return m
}
