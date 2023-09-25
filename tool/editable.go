package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/andrewarrow/feedback/router"
)

func editable(path, name string) {
	asBytes, _ := ioutil.ReadFile(path + "/app/feedback.json")
	var site router.FeedbackSite
	json.Unmarshal(asBytes, &site)
	m := site.FindModel(name)

	fmt.Println(`<table class="inline-block whitespace-nowrap">`)
	fmt.Println(`{{$row := index . "item"}}`)
	for _, field := range m.Fields {
		fmt.Println("<tr>")
		fmt.Println(`{{$value := index $row "` + field.Name + `"}}`)
		t := `
<td class="px-3 py-2">
  %s<br/>
  {{ textfield "%s" $value}}
</td>
`
		fmt.Printf(t, field.Name, field.Name)
		fmt.Println("</tr>")
	}
	fmt.Println("</table>")
}
