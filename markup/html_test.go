package markup

import (
	"fmt"
	"testing"
)

func TestToHTML(t *testing.T) {
	send := map[string]any{}
	lines := []string{
		"div",
		"  div",
		"    div"}

	s := ToHTMLFromLines(send, lines)
	fmt.Println(s)
}
