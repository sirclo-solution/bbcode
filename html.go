package bbcode

import (
	"github.com/sirclo-solution/bbcode/bbnode"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"strings"
)

const (
	newLine = "\n"
)

func HTMLToBBCode(r io.Reader) (res string, err error) {
	bb, err := HTMLToBBNode(r)
	return bb.String(), err
}

func trimHTML(data string) string {
	tmp := data
	tmp = strings.Replace(tmp, "\r", "", -1)
	tmp = strings.Replace(tmp, "\n", "", -1)
	tmp = strings.Replace(tmp, "	", "", -1)
	return tmp
}

func HTMLToBBNode(r io.Reader) (res bbnode.Interface, err error) {
	doc, err := html.Parse(r)
	if err == nil {
		res = HTMLNodeToBBNode(doc, nil)
	}
	return res, err
}

func HTMLNodeToBBNode(node *html.Node, parent bbnode.Interface) (res bbnode.Interface) {
	if node != nil {
		switch node.Type {
		case html.ElementNode:
			switch node.DataAtom {
			case atom.Html, atom.Tbody:
				//ignore tag
				res = new(bbnode.Node)
			case atom.Head, atom.Script, atom.Style:
				//ignore all
				res = nil
			case atom.Body:
				res = new(bbnode.Node)
			case atom.Br:
				res = bbnode.NewBrNode()
			case atom.P:
				res = bbnode.NewPNode()
			case atom.B, atom.Strong:
				res = bbnode.NewBBNode("b")
			case atom.I, atom.Em:
				res = bbnode.NewBBNode("i")
			case atom.U, atom.Ins:
				res = bbnode.NewBBNode("u")
			case atom.S, atom.Del:
				res = bbnode.NewBBNode("s")
			case atom.Pre:
				res = bbnode.NewBlockNode("code")
			case atom.Blockquote:
				res = bbnode.NewBlockNode("quote")
			case atom.Img:
				var url string
				for _, a := range node.Attr {
					if a.Key == "src" {
						url = a.Val
						break
					}
				}
				res = bbnode.NewBBNode("img")

				urlNode := new(bbnode.Node)
				urlNode.Value = url
				res.AddChild(urlNode)
			case atom.A:
				var url string
				for _, a := range node.Attr {
					if a.Key == "href" {
						url = a.Val
						break
					}
				}
				res = bbnode.NewBBNode("url", url)
			case atom.Ul, atom.Ol:
				res = bbnode.NewBlockNode("list")
			case atom.Li:
				res = bbnode.NewLiNode()
			case atom.Table, atom.Tr:
				res = bbnode.NewBlockNode(node.Data)
			case atom.Td:
				res = bbnode.NewLineNode(node.Data)
			default:
				res = new(bbnode.Node)
			}
		default:
			res = bbnode.NewNode(trimHTML(node.Data))
		}

		if res != nil {
			res.SetParent(parent)
			for cHTML := node.FirstChild; cHTML != nil; cHTML = cHTML.NextSibling {
				cBB := HTMLNodeToBBNode(cHTML, res)
				res.AddChild(cBB)
			}
		}
	}
	return res
}
