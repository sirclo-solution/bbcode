package bbnode

import (
	"fmt"
	"strings"
)

type BlockNode struct {
	BBNode
}

func (node *BlockNode) String() string {
	child := node.sChild()
	if !strings.HasSuffix(child, "\n") {
		child += "\n"
	}
	return fmt.Sprintf("%s\n%s%s\n", node.sOpenTag(), child, node.sCloseTag())
}

func (node *BlockNode) SHTML() string {
	child := node.SHTMLChild()
	tag := strings.ToLower(node.Data)
	if tag == "code" {
		tag = "pre"
	} else if tag == "quote" {
		tag = "blockquote"
	} else if tag == "list" {
		tag = "ul"
	}
	return fmt.Sprintf("<%s>\n%s</%s>\n", tag, child, tag)
}

func (node *BlockNode) Type() string {
	return "block"
}

func NewBlockNode(data string, attr ...string) *BlockNode {
	res := BlockNode{}
	res.Data = data

	if len(attr) > 0 {
		//first attr set to value
		res.Value = attr[0]
		res.Attr = attr[1:len(attr)]
	}
	return &res
}
