package util

import (
	"os/exec"
	"strings"
)

func PipeToJq(inputString string) string {

	jq := exec.Command("jq", ".")
	jq.Stdin = strings.NewReader(inputString)

	b, _ := jq.CombinedOutput()
	return string(b)

}
