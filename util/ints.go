package util

import (
	"fmt"
	"strings"
)

func IntsToStringList(a []int64) string {
	buffer := []string{}
	for _, item := range a {
		s := fmt.Sprintf("%d", item)
		buffer = append(buffer, s)
	}
	return strings.Join(buffer, ",")
}
