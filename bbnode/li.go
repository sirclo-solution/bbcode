package bbnode

import (
	"fmt"
)

type LiNode struct {
	BBNode
}

func (node *LiNode) String() string {
	return fmt.Sprintf("[*]%s\n", node.sChild())
}

func (node *LiNode) SHTML() string {
	return fmt.Sprintf("<li>%s</li>\n", node.SHTMLChild())
}

func NewLiNode() *LiNode {
	res := LiNode{}
	return &res
}
