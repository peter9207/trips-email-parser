package html

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type Element struct {
	Type     string
	Body     string
	Props    []string
	Children []Element
}

type Keyword struct {
	Word   string
	Format string
}

var keywords = []Keyword{
	{
		Word:   "CHECK-IN",
		Format: "\\d{1,2}\\/\\d{1,2}\\/\\d{2,4}",
	},
}

var resultMap map[string]string

func PrintNode(prefix string, n *html.Node, newLine bool) {

	var t string
	switch n.Type {
	case html.TextNode:
		t = "Text"
	case html.DocumentNode:
		t = "Doc"
	case html.ElementNode:
		t = "Ele"
	case html.CommentNode:
		t = "Comment"
	case html.DoctypeNode:
		t = "Doctype"
	}

	trimmed := strings.TrimSpace(n.Data)
	fmt.Printf("<%s %v %s %s>", prefix, t, n.Namespace, trimmed)
	if newLine {
		fmt.Println()
	}
}

func ContainsDate(data string) bool {

	var dateRegex = []string{
		"\\d{4}-\\d{2}-\\d{2}",
	}

	for _, pattern := range dateRegex {
		if ok, _ := regexp.MatchString(pattern, data); ok {
			return true
		}

	}
	return false
}

func ContainsKeyWords(data string) bool {

	var keywords = []string{
		"CHECK-IN",
	}

	for _, keyword := range keywords {
		if strings.Contains(data, keyword) {
			fmt.Println("found keyword", data)
			return true
		}
	}
	return false
}

func PrintTree(root *html.Node) {
	// PrintNode("root", root, true)

	child := root.FirstChild

	for child != nil {
		PrintTree(child)
		data := strings.ToUpper(child.Data)

		if ContainsDate(data) {
			fmt.Println("found date", data)

		}
		// if ContainsKeyWords(data) {
		// 	fmt.Println(data)

		// 	PrintNode("found", child, true)
		// }

		child = child.NextSibling
	}

}

func Treefy(reader io.Reader) (err error) {

	doc, err := html.Parse(reader)
	if err != nil {
		return
	}

	PrintTree(doc)

	// fmt.Printf("parentNode %v, %v", )
	// PrintNode("root", doc)

	// child := doc.FirstChild

	// for child != nil {
	// 	PrintNode("child", child)
	// 	child = child.NextSibling
	// }
	return
}
