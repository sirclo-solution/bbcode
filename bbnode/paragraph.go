package bbnode

type PNode struct {
	Node
}

func (node *PNode) String() string {
	return "\n" + node.Node.String() + "\n\n"
}

func NewPNode() *PNode {
	res := PNode{}
	return &res
}
