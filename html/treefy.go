package html

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Element struct {
	Type     string
	Body     string
	Props    []string
	Children []Element
}

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

func PrintTree(root *html.Node) {
	PrintNode("root", root, true)

	child := root.FirstChild

	for child != nil {
		PrintTree(child)
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
