package markup

import (
	"fmt"
	"strings"
)

type Tag struct {
	Name     string
	Text     string
	Children []*Tag
	Close    bool
	Attr     map[string]string
}

var validTagMap = map[string]int{"div": 2, "img": 3, "root": 1, "a": 2,
	"iframe": 2, "p": 2, "span": 2, "form": 2, "input": 3, "textarea": 2,
	"button": 2, "{{": 4}

func NewTag(index int, tokens []string) *Tag {
	t := Tag{}
	name := tokens[index]
	t.Attr = makeClassAndAttrMap(name, tokens[index+1:len(tokens)])
	//if name == "img" {
	//	t.Attr["class"] += "w-20 "
	//}
	flavor := validTagMap[name]
	t.Name = name
	if flavor > 0 && flavor < 4 {
		t.Close = flavor == 2
	} else {
		t.Text = strings.Join(tokens[index:len(tokens)], " ")
	}
	if flavor == 0 {
		t.Name = ""
	}
	t.Children = []*Tag{}
	//t.Parent = parent
	return &t
}

func (t *Tag) MakeAttr() string {
	buffer := ""

	for key, value := range t.Attr {
		buffer += fmt.Sprintf(`%s="%s" `, key, value)
	}

	return buffer
}

func fixValueForTag(name, key, value string) string {
	if strings.HasPrefix(value, "http") {
		return value
	}
	if name == "img" && key == "src" {
		value = fmt.Sprintf("/assets/images/%s", value)
	}
	return value
}

func getKeyValue(s string) (string, string) {
	tokens := strings.Split(s, "=")
	if len(tokens) == 2 {
		return tokens[0], tokens[1]
	}
	return "", ""
}

func makeClassAndAttrMap(name string, tokens []string) map[string]string {
	m := map[string]string{}

	class := ""
	for _, item := range tokens {
		if strings.Contains(item, "=") {
			key, value := getKeyValue(item)
			value = fixValueForTag(name, key, value)
			m[key] = value
		} else {
			if item == "bg-r" {
				item = randomColor()
			}
			class += item + " "
		}
	}
	m["class"] = class

	return m
}
