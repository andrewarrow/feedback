package markup

import (
	"fmt"
	"strings"
	"testing"
)

func TestToHTML(t *testing.T) {
	send := map[string]any{}

	q := `
div
  div
    div bg-green-100 space-y-3 pt-3 pl-3
      {{ $list := index . "list" }}
      {{ range $i, $item := $list }}
        div flex 
          div mr-3
            {{ add $i 1 }}.
          div
            {{ $title := index $item "title" }}
            {{ $id := index $item "id_hacker" }}
            {{ $digitSum := index $item "digit_sum" }}
            {{ $sum := index $item "sum" }}
            div
              a href=https://news.ycombinator.com/item?id!{{$id}}
                {{ $title }}
`
	lines := strings.Split(q, "\n")
	s := ToHTMLFromLines(send, lines)
	fmt.Println(s)
}
