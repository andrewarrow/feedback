package server

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/justincampbell/timeago"
	"github.com/russross/blackfriday"
)

func AddTemplates(r *gin.Engine) {
	fm := template.FuncMap{
		"mod": func(i, j int) bool { return i%j == 0 },
		"ago": func(i int64) string {
			d, _ := time.ParseDuration(fmt.Sprintf("%ds", time.Now().Unix()-i))
			return timeago.FromDuration(d)
		},
		"tokens": func(s string, i int) string { return strings.Split(s, ".")[i] },
		"add":    func(i, j int) int { return i + j },
		"md": func(s string) template.HTML {
			md := blackfriday.MarkdownCommon([]byte(s))
			return template.HTML(string(md))
		},
	}
	r.SetFuncMap(fm)
	r.LoadHTMLGlob("templates/*.tmpl")
}
