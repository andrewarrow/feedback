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
			  hi4
      div 5 
        div 6
          div 7
            div 8
						  hi8
          div 9
						hi9
`
	lines := strings.Split(q, "\n")
	s := ToHTMLFromLines(send, lines)
	fmt.Println(s)
}
