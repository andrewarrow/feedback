package server

import "github.com/gin-gonic/gin"

import "fmt"
import "time"
import "strings"
import "github.com/russross/blackfriday"
import "github.com/justincampbell/timeago"
import "html/template"

func AddTemplates(r *gin.Engine, prefix string) {
	fm := template.FuncMap{
		"mod": func(i, j int) bool { return i%j == 0 },
		"ago": func(i int64) string {
			d, _ := time.ParseDuration(fmt.Sprintf("%ds", time.Now().Unix()-i))
			return timeago.FromDuration(d)
		},
		"tokens": func(s string, i int) string { return strings.Split(s, ".")[i] },
		"add":    func(i, j int) int { return i + j },
		"md": func(s string) template.HTML {
			md := blackfriday.Run([]byte(s))
			return template.HTML(string(md))
		},
	}
	r.SetFuncMap(fm)
	r.LoadHTMLGlob(prefix + "templates/*.tmpl")
}
