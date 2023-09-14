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
	//m := site.FindModel(name)

	fmt.Println("func handleFoo(c *router.Context) {")
	fmt.Printf(`list := c.All("%s", "order by created_at desc", "")\n`, name)
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

	buff := []string{}
	for _, field := range m.Fields {
		buff = append(buff, "<th>"+field.Name+"</th>")
	}
	header := strings.Join(buff, "\n")

	fmt.Println("<table>")
	fmt.Println(header)
	fmt.Println(`{{$list := index . "list"}}`)
	fmt.Println(`{{range $i, $row := $list}}`)
	for _, field := range m.Fields {
		fmt.Println(`{{$thing := index $row "` + field.Name + `"}}`)
		fmt.Println("<td>{{$thing}}</td>")
	}
	fmt.Println(`{{end}}`)
	fmt.Println("</table>")
}
