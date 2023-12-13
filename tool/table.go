package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/andrewarrow/feedback/router"
)

func list(path, name string) {
	asBytes, _ := ioutil.ReadFile(path + "/app/feedback.json")
	var site router.FeedbackSite
	json.Unmarshal(asBytes, &site)
	m := site.FindModel(name)
	m.EnsureIdAndCreatedAt()

	top := `package app

import (
  "github.com/andrewarrow/feedback/router"
)`

	fmt.Println(top)
	fmt.Println("func handleFoo(c *router.Context) {")
	fmt.Printf(`list := c.All("%s", "order by created_at desc", "")`+"\n", name)
	send := `send := map[string]any{}
	send["list"] = list
	c.SendContentInLayout("foo.html", send, 200)`
	fmt.Println(send)
	fmt.Println("}")
}

func table(path, name string) {
	asBytes, _ := ioutil.ReadFile(path + "/app/feedback.json")
	var site router.FeedbackSite
	json.Unmarshal(asBytes, &site)
	m := site.FindModel(name)
	m.EnsureIdAndCreatedAt()

	buff := []string{}
	goal := `  {{ $items := index . "items" }}
  table
    tr font-bold`
	buff = append(buff, goal)
	for _, field := range m.Fields {
		buff = append(buff, "      td\n        "+field.Name)
	}
	header := strings.Join(buff, "\n")
	fmt.Println(header)

	fmt.Println(`    {{ range $i, $item := $items }}`)
	fmt.Println(`      tr`)
	for _, field := range m.Fields {
		fmt.Println(`      {{ $` + field.Name + ` := index $item "` + field.Name + `" }}`)
		fmt.Printf(`        td` + "\n          {{ $" + field.Name + " }}\n")
	}
	fmt.Println(`    {{end}}`)
}
