package router

import (
	"html/template"
	"strings"
	"time"

	"github.com/xeonx/timeago"
)

func TemplateFunctions() template.FuncMap {
	fm := template.FuncMap{
		"mod":    func(i, j int) bool { return i%j == 0 },
		"tokens": func(s string, i int) string { return strings.Split(s, ".")[i] },
		"add":    func(i, j int) int { return i + j },
		"ago":    func(t time.Time) string { return timeago.English.Format(t) },
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
	t, _ = t.ParseGlob("views/*.html")
	return t
}
