package markup

import (
	"fmt"
	"strings"
	"testing"
)

func TestToHTML(t *testing.T) {
	send := map[string]any{}

	q := `
div 1
  div 2
    div 3
      div 4
        div 5
  div 6
    hi
`
	lines := strings.Split(q, "\n")
	s := ToHTMLFromLines(send, lines)
	fmt.Println(s)
}
