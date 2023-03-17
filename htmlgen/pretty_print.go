package htmlgen

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/andrewarrow/feedback/files"
	"golang.org/x/net/html"
)

func PrettyPrint() {
	list, _ := ioutil.ReadDir("views")
	for _, file := range list {
		input := files.ReadFile("views/" + file.Name())
		s := parseIt(input)
		fmt.Println(s)
	}
}

func parseIt(input string) string {
	doc, err := html.ParseFragment(strings.NewReader(input), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing HTML: %v\n", err)
		os.Exit(1)
	}
	content := new(bytes.Buffer)
	for _, n := range doc {
		indent(content, n, 0)
	}

	return content.String()
}

func indent(w io.Writer, n *html.Node, depth int) {
	prefix := strings.Repeat(" ", depth*2)
	switch n.Type {
	case html.ElementNode:
		fmt.Fprintf(w, "%s<%s", prefix, n.Data)
		for _, a := range n.Attr {
			fmt.Fprintf(w, " %s=\"%s\"", a.Key, a.Val)
		}
		if n.FirstChild == nil {
			fmt.Fprint(w, "/>\n")
		} else {
			fmt.Fprint(w, ">\n")
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				indent(w, c, depth+1)
			}
			fmt.Fprintf(w, "%s</%s>\n", prefix, n.Data)
		}
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Fprintf(w, "%s%s\n", prefix, text)
		}
	}
}
