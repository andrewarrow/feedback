package aigen

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/andrewarrow/feedback/network"
)

func RunImage(prompt string) {

	m := map[string]any{"prompt": "work from home",
		"n":    1,
		"size": "256x256"}
	asBytes, _ := json.Marshal(m)

	key := os.Getenv("OPEN_AI")
	fmt.Println(key)

	jsonString, _ := network.DoPost(nil, os.Getenv("OPEN_AI"), "/v1/images/generations", asBytes)
	fmt.Println(jsonString)

}
