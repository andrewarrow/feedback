package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/andrewarrow/feedback/router"
)

func table(path, name string) {
	asBytes, _ := ioutil.ReadFile(path + "/app/feedback.json")
	var site router.FeedbackSite
	json.Unmarshal(asBytes, &site)
	m := site.FindModel(name)

	for _, field := range m.Fields {
		fmt.Println(field.Name)
	}

}
