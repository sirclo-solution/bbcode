package bbnode

import (
	"fmt"
	"strings"
)

type BBNode struct {
	Node
	Data string
	Attr []string
}

func (node *BBNode) String() string {
	return node.sOpenTag() + node.sChild() + node.sCloseTag()
}

func (node *BBNode) sOpenTag() string {
	var value, attr string
	if node.Value != "" {
		value = "=" + node.Value
	}
	if len(node.Attr) > 0 {
		attr = " " + strings.Join(node.Attr, " ")
	}
	return fmt.Sprintf("[%s%s%s]", node.Data, value, attr)
}

func (node *BBNode) sCloseTag() string {

	return fmt.Sprintf("[/%s]", node.Data)
}

func (node *BBNode) sChild() string {
	var res string
	for c := node.FirstChild(); c != nil; c = c.NextSibling() {
		res += c.String()
	}
	return res
}

func (node *BBNode) SHTML() (res string) {
	child := node.SHTMLChild()
	tag := strings.ToLower(node.Data)
	if tag == "url" {
		url := node.Value
		if url == "" {
			url = child
		}
		res = fmt.Sprintf("<a href=\"%s\">%s</a>", node.Value, child)
	} else if tag == "img" {
		res = fmt.Sprintf("<%s src=\"%s\"/>", tag, child)
	} else {
		res = fmt.Sprintf("<%s>%s</%s>", tag, child, tag)
	}
	return res
}

func (node *BBNode) SHTMLChild() string {
	var res string
	for c := node.FirstChild(); c != nil; c = c.NextSibling() {
		res += c.SHTML()
	}
	return res
}

//NewBBNode construct new BBNode, second argument will be set to Value, the remaining set to Attr
func NewBBNode(data string, attr ...string) *BBNode {
	res := BBNode{
		Data: data,
	}
	if len(attr) > 0 {
		//first attr set to value
		res.Value = attr[0]
		res.Attr = attr[1:len(attr)]
	}
	return &res
}
