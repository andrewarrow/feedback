package util

import (
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	s := strings.ToLower(snake)
	return s
}
func ToCamelCase(str string) string {
	return strcase.ToCamel(str)
}

func FixForDash(s string) string {
	return strings.Replace(s, "-", "_", -1)
}
