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
      {{ test }}
      {{ range }}
        div 5
        div 55
          hi
      {{ end }}
  div 6
    hi
`
	lines := strings.Split(q, "\n")
	s := ToHTMLFromLines(send, lines)
	fmt.Println(s)
}
