package router

import (
	"embed"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/xeonx/timeago"
)

var EmbeddedTemplates embed.FS

func TemplateFunctions() template.FuncMap {
	fm := template.FuncMap{
		"mod":    func(i, j int) bool { return i%j == 0 },
		"tokens": func(s string, i int) string { return strings.Split(s, ".")[i] },
		"add":    func(i, j int) int { return i + j },
		"null": func(s any) any {
			if s == nil {
				return ""
			}
			return s
		},
		"ago": func(t time.Time) string { return timeago.English.Format(t) },
		"adds": func(i int64) string {
			if i != 1 {
				return "s"
			}
			return ""
		},
	}
	return fm
}

func LoadTemplates() *template.Template {
	t := template.New("")
	t = t.Funcs(TemplateFunctions())

	templateFiles, _ := EmbeddedTemplates.ReadDir("views")
	for _, file := range templateFiles {
		fileContents, _ := EmbeddedTemplates.ReadFile("views/" + file.Name())
		_, err := t.New(file.Name()).Parse(string(fileContents))
		if err != nil {
			fmt.Println(file.Name(), err)
		}
	}
	return t
}
