package markup

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ToHTML(m map[string]any, filename string) string {
	asBytes, _ := ioutil.ReadFile("markup/" + filename)
	asString := string(asBytes)
	asLines := strings.Split(asString, "\n")
	root := NewTag(0, []string{"root"})

	stack := []*Tag{root}
	var lastSpaces int

	for _, line := range asLines {
		tokens := strings.Split(line, " ")
		if len(tokens) == 1 {
			continue
		}

		spaces := countSpaces(tokens)
		delta := spaces - lastSpaces
		//fmt.Println(delta, line)
		if delta < 0 {
			delta = delta * -1
			delta = delta / 2
			offset := 2
			if delta > 2 {
				offset = (delta * 2) - 2
			}
			//fmt.Println("f", delta, offset, line)
			stack = stack[0 : len(stack)-(offset)]
		}
		if spaces == 0 {
			stack = []*Tag{root}
		}
		if spaces == 4 && len(stack) == 4 {
			stack = stack[0:3]
		}

		tag := NewTag(spaces, tokens)
		parent := stack[len(stack)-1]
		parent.Children = append(parent.Children, tag)
		stack = append(stack, tag)

		lastSpaces = spaces
	}

	final := renderHTML(m, root, "")
	fmt.Println(final)
	return final
}

func renderHTML2(m map[string]any, tag *Tag, tabs string) string {
	if tag.Name != "root" && tag.Name != "" {
		fmt.Println(tabs + tag.Name)
	}

	for _, child := range tag.Children {
		renderHTML(m, child, tabs+"  ")
	}

	if tag.Name != "root" && tag.Name != "" && tag.Close {
		fmt.Println(tabs + "/" + tag.Name)
	}

	if tag.Text != "" {
		fmt.Println(tabs + "/" + tag.Text)
	}

	return ""
}

func renderHTML(m map[string]any, tag *Tag, tabs string) string {
	html := ""

	if tag.Name != "root" && tag.Name != "" {
		html += fmt.Sprintf("%s<%s", tabs, tag.Name)
		//html += tabs + "<" + tag.Name
		html += fmt.Sprintf(` %s `, tag.MakeAttr())
		if tag.Close == false {
			html += "/>"
		} else {
			html = strings.TrimRight(html, " ") + ">"
		}
		html += "\n"
	}

	for _, child := range tag.Children {
		html += renderHTML(m, child, tabs+"  ")
	}

	if tag.Name != "root" && tag.Name != "" && tag.Close {
		html += tabs + "</" + tag.Name + ">"
		html += "\n"
	}

	if tag.Text != "" {
		if strings.HasPrefix(tag.Text, "#") {
			key := tag.Text[1:len(tag.Text)]
			html += m[key].(string)
		} else {
			html += tabs + tag.Text
			html += "\n"
		}
	}

	return html
}

func countSpaces(tokens []string) int {
	count := 0
	for _, item := range tokens {
		if item == "" {
			count++
		} else {
			break
		}
	}
	return count
}
