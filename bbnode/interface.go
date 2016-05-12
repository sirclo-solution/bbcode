package bbnode

type Interface interface {
	String() string
	SHTML() string
	Type() string
	Parent() Interface
	FirstChild() Interface
	LastChild() Interface
	PrevSibling() Interface
	NextSibling() Interface
	SetParent(Interface)
	SetFirstChild(Interface)
	SetLastChild(Interface)
	AddChild(Interface)
	SetPrevSibling(Interface)
	SetNextSibling(Interface)
}

type Node struct {
	parent, firstChild, lastChild, prevSibling, nextSibling Interface

	Value string
}

func (node *Node) String() string {
	return node.Value + node.SChild()
}

func (node *Node) SChild() string {
	var child string
	for c := node.FirstChild(); c != nil; c = c.NextSibling() {
		child += c.String()
	}

	return child
}

func (node *Node) SHTML() string {
	return node.Value + node.SHTMLChild()
}

func (node *Node) SHTMLChild() string {
	var child string
	for c := node.FirstChild(); c != nil; c = c.NextSibling() {
		child += c.SHTML()
	}

	return child
}

func (node *Node) Type() string {
	return "node"
}

func (node *Node) Parent() Interface {
	return node.parent
}

func (node *Node) FirstChild() Interface {
	return node.firstChild
}

func (node *Node) LastChild() Interface {
	return node.lastChild
}

func (node *Node) PrevSibling() Interface {
	return node.prevSibling
}

func (node *Node) NextSibling() Interface {
	return node.nextSibling
}

func (node *Node) SetParent(n Interface) {
	node.parent = n
}

func (node *Node) SetFirstChild(n Interface) {
	node.firstChild = n
}

func (node *Node) SetLastChild(n Interface) {
	node.lastChild = n
}

func (node *Node) AddChild(n Interface) {
	if node.firstChild == nil {
		node.firstChild = n

	}
	if node.lastChild != nil {
		last := node.lastChild
		last.SetNextSibling(n)
		n.SetPrevSibling(last)
	}
	node.lastChild = n
	if n != nil {
		n.SetParent(node)
	}
}

func (node *Node) SetPrevSibling(n Interface) {
	node.prevSibling = n
}

func (node *Node) SetNextSibling(n Interface) {
	node.nextSibling = n
}

func NewNode(value string) *Node {
	res := Node{
		Value: value,
	}
	return &res
}
