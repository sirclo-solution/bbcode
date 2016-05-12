package bbnode

import (
	"fmt"
)

type BrNode struct {
	Node
}

func (node *BrNode) String() string {
	return "\n"
}

func (node *BrNode) SHTML() string {
	return "<br/>\n"
}

func NewBrNode() *BrNode {
	res := BrNode{}
	return &res
}

type LineNode struct {
	BBNode
}

func (node *LineNode) String() string {
	return fmt.Sprintf("%s%s%s\n", node.sOpenTag(), node.sChild(), node.sCloseTag())
}

func (node *LineNode) SHTML() string {
	return fmt.Sprintf("<%s>%s</%s>\n", node.Data, node.SHTMLChild(), node.Data)
}

func (node *LineNode) Type() string {
	return "line"
}

func NewLineNode(data string, attr ...string) *LineNode {
	res := LineNode{}
	res.Data = data

	if len(attr) > 0 {
		//first attr set to value
		res.Value = attr[0]
		res.Attr = attr[1:len(attr)]
	}
	return &res
}
