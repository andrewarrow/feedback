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

const DATE_LAYOUT = "Monday, January 2, 2006 15:04"

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
		"epoch": func(date string, hour float64, tzString string) int64 {
			dateString := fmt.Sprintf("%s %02d:00", date, int(hour))
			tz, _ := time.LoadLocation(tzString)
			t, _ := time.ParseInLocation(DATE_LAYOUT, dateString, tz)
			utc, _ := time.LoadLocation("UTC")
			return t.In(utc).Unix()
		},
		"price": func(pennies int64) string {
			amount := fmt.Sprintf("%d", pennies)
			if len(amount) < 3 {
				return fmt.Sprintf("$00.%s USD", amount)
			}
			dollars := amount[0 : len(amount)-2]
			cents := amount[len(amount)-2:]
			return fmt.Sprintf("$%s.%s USD", dollars, cents)
		},
		"ampm": func(f float64) string {
			i := int(f)
			if i > 12 {
				return fmt.Sprintf("%02d:00 PM", i-12)
			}
			if i == 0 {
				return "12:00 AM"
			}
			return fmt.Sprintf("%02d:00 AM", i)
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
