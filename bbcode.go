package bbcode

import (
	"github.com/sirclo-solution/bbcode/bbnode"
	"strings"
)

func BBCodeToHTML(input string) (res string, err error) {
	node := Parse(input, "")
	Simplify(node)
	root := new(bbnode.Node)
	BBfy(root, node)
	return root.SHTML(), err
}

func BBfy(root bbnode.Interface, cur *bbnode.Node) (res bbnode.Interface) {
	if cur != nil {
		tag := strings.ToLower(cur.Value)
		tag = strings.TrimPrefix(tag, "[")
		tag = strings.TrimSuffix(tag, "]")
		switch tag {
		case "b", "i", "u", "s", "img":
			res = bbnode.NewBBNode(tag)
		case "code", "quote", "list", "table", "tr":
			res = bbnode.NewBlockNode(tag)
		case "*":
			res = bbnode.NewLiNode()
		case "td":
			res = bbnode.NewLineNode(tag)
		default:
			if strings.HasPrefix(tag, "url") {
				var url string
				eq := strings.Index(tag, "=")
				if eq != -1 {
					space := strings.Index(tag[eq+1:len(tag)], " ")
					if space != -1 {
						url = tag[eq+1 : space]
					} else {
						url = tag[eq+1 : len(tag)]
					}
				}
				if url == "" {
					url = cur.SChild()
				}
				res = bbnode.NewBBNode("url", url)
			} else if cur.Value == "\n" {
				res = bbnode.NewBrNode()
			} else {
				res = bbnode.NewNode(cur.Value)
			}
		}

		if res != nil {
			root.AddChild(res)
			if cur.FirstChild() != nil {
				child := cur.FirstChild()
				if res.Type() == "block" {
					if child.(*bbnode.Node).Value == "\n" {
						child = child.NextSibling()
					}
				}
				if child != nil {
					BBfy(res, child.(*bbnode.Node))
				}
			}
		}
		if cur.NextSibling() != nil {
			sibling := cur.NextSibling()
			if res != nil {
				if res.Type() == "block" || res.Type() == "line" {
					if sibling.(*bbnode.Node).Value == "\n" {
						sibling = sibling.NextSibling()
					}
				}
			}
			if sibling != nil {
				BBfy(root, sibling.(*bbnode.Node))
			}
		}
	}

	return
}

func Simplify(root *bbnode.Node) *bbnode.Node {
	// var stack []*bbnode.Node

	var parent *bbnode.Node
	parent = nil
	for cur := root; cur != nil; {
		cur.SetParent(parent)

		if (strings.HasPrefix(cur.Value, "[/") || cur.Value == "\n") && (parent != nil) {
			var prefix string
			if cur.Value == "\n" {
				prefix = "[*]"
			} else {
				prefix = strings.TrimSuffix(strings.Replace(cur.Value, "/", "", 1), "]")
			}

			for tmp := cur.Parent(); tmp != nil; tmp = tmp.Parent() {
				node := tmp.(*bbnode.Node)
				if node != nil {
					if strings.HasPrefix(node.Value, prefix) {
						if cur.NextSibling() != nil {
							next := cur.NextSibling()
							node.SetNextSibling(next)
							next.SetPrevSibling(node)

							if cur.PrevSibling() != nil {
								prev := cur.PrevSibling()
								prev.SetNextSibling(nil)
							}
							cur = node
						} else {
							cur = nil
						}

						if node.Parent() != nil {
							parent = node.Parent().(*bbnode.Node)
						} else {
							parent = nil
						}
						break
					}
					nodeParent := node.Parent().(*bbnode.Node)
					if nodeParent == nil {
						break
					}
				}
			}
		} else if strings.HasPrefix(cur.Value, "[") {
			if cur.NextSibling() != nil {
				parent = cur
				next := cur.NextSibling().(*bbnode.Node)
				cur.AddChild(next)
				cur.SetNextSibling(nil)
				next.SetPrevSibling(nil)
				cur = next
			} else {
				cur = nil
			}
			continue
		}

		if cur.NextSibling() != nil {
			cur = cur.NextSibling().(*bbnode.Node)
		} else {
			cur = nil
		}
	}
	return root
}

func NewSibling(node *bbnode.Node, t string) *bbnode.Node {
	res := bbnode.NewNode(t)
	res.SetParent(node.Parent())
	node.SetNextSibling(res)
	res.SetPrevSibling(node)
	return res
}

func Parse(input string, mode string) *bbnode.Node {
	root := bbnode.Node{}

	cur := &root

	for i := range input {
		inp := string(input[i])
		switch inp {
		case "[":
			if cur.Value == "" {
				cur.Value += string(input[i])
			} else {
				cur = NewSibling(cur, "[")
			}
		case "]":
			if strings.HasPrefix(cur.Value, "[") {
				cur.Value += inp
				cur = NewSibling(cur, "")
			} else if cur.Parent != nil && strings.HasPrefix(cur.Parent().(*bbnode.Node).Value, "[") {
				cur = cur.Parent().(*bbnode.Node)
				cur.Value += string(input[i])
			} else {
				cur.Value += string(input[i])
			}
		case "\n":
			if cur.Value == "" {
				cur.Value += "\n"
			} else {
				cur = NewSibling(cur, "\n")
			}
			cur = NewSibling(cur, "")
		default:
			cur.Value += inp
		}
	}

	return &root
}
