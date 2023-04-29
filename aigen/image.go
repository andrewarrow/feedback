package aigen

import (
	"encoding/json"
	"fmt"

	"github.com/andrewarrow/feedback/network"
)

func RunImage(prompt string) {

	m := map[string]any{"prompt": "a logo like the better business bureau but called \"remote renters\" that tells consumers this apartment building does specific stuff for people that work from home",
		"n":    2,
		"size": "1024x1024"}
	m = map[string]any{"prompt": "work from home",
		"n":    2,
		"size": "1024x1024"}
	asBytes, _ := json.Marshal(m)

	jsonString, _ := network.DoPost(nil, "", "/v1/image/generations", asBytes)
	fmt.Println(jsonString)

}
