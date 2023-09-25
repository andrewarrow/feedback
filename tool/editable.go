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
		if field.Flavor == "photo" {
			t = `
<td class="px-3 py-2">
  %s<br/>
    {{ if eq $value ""}}
    <img src="https://i.imgur.com/OY8fov3.png" width="160"/>
    {{ else }}
    <img src="/bucket/{{$value}}" width="160"/>
    {{ end }}
   <br/>
   <input type="file" name="file" accept="image/jpeg, image/png, image/gif"/>
</td>
`
		}
		fmt.Printf(t, field.Name, field.Name)
		fmt.Println("</tr>")
	}
	fmt.Println("</table>")
}
