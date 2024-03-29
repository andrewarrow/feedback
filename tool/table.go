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
	buff = append(buff, fmt.Sprintf(`{{ define "%ss" }}`, name))
	goal := `  {{ $items := index . "items" }}
  table id=thing
    tr font-bold`
	buff = append(buff, goal)
	for _, field := range m.Fields {
		buff = append(buff, "      td pr-3\n        "+field.Name)
	}
	header := strings.Join(buff, "\n")
	fmt.Println(header)

	fmt.Println(`    {{ range $i, $item := $items }}`)
	fmt.Println(`      tr`)
	fmt.Printf(`        {{ template "%s" $item }}`+"\n", name)
	fmt.Println(`    {{end}}`)
	fmt.Println(`  {{end}}`)

	fmt.Printf(`  {{ define "%s" }}`+"\n", name)
	for _, field := range m.Fields {
		fmt.Println(`  {{ $` + field.Name + ` := index . "` + field.Name + `" }}`)
		fmt.Printf(`    td pr-3 whitespace-nowrap` + "\n      {{ $" + field.Name + " }}\n")
	}
	fmt.Println(`  {{end}}`)

}
