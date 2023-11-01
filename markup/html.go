package markup

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ToHTML(m map[string]any, filename string) string {
	asBytes, _ := ioutil.ReadFile(filename)
	asString := string(asBytes)
	asLines := strings.Split(asString, "\n")
	return ToHTMLFromLines(m, asLines)
}

func ToHTMLFromLines(m map[string]any, asLines []string) string {
	root := NewTag(0, []string{"root"})

	spaceMap := map[string]*Tag{}
	for i, line := range asLines {
		tokens := strings.Split(line, " ")
		if len(tokens) == 1 {
			continue
		}
		spaces := countSpaces(tokens)
		tag := NewTag(spaces, tokens)
		key := fmt.Sprintf("%d_%d", i, spaces)
		//fmt.Println(key)
		spaceMap[key] = tag
	}

	//fmt.Println("key")

	for i, line := range asLines {
		tokens := strings.Split(line, " ")
		if len(tokens) == 1 {
			continue
		}

		spaces := countSpaces(tokens)
		more := 0
		for {
			key := fmt.Sprintf("%d_%d", i-more, spaces-2)
			p := spaceMap[key]
			if p != nil {
				key = fmt.Sprintf("%d_%d", i, spaces)
				t := spaceMap[key]
				//fmt.Println(key, p.Name, t.Name)

				p.Children = append(p.Children, t)
				break
			}
			more++
			if i-more < 0 {
				break
			}
		}

	}

	top := spaceMap["0_0"]
	root.Children = append(root.Children, top)
	final := renderHTML(m, root, "")
	//final := ""
	//fmt.Println(root)
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
		if tag.Name != "{{" {
			html += fmt.Sprintf("%s<%s", tabs, tag.Name)
			//html += tabs + "<" + tag.Name
			html += fmt.Sprintf(` %s `, tag.MakeAttr())
		}
		if tag.Close == false {
			if tag.Name != "{{" {
				html += "/>"
			} else {
				html += fmt.Sprintf("%s%s", tabs, tag.Text)
			}
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

	if tag.Text != "" && tag.Name != "{{" {
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
