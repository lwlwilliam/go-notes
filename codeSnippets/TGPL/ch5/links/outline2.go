package links

import (
	"golang.org/x/net/html"
	"fmt"
)

var depth int

// 针对每个结点 x，都会调用 pre(x) 和 post(x)
// pre, post 都是可选的
// 遍历孩子结点之前，pre 被调用
// 遍历孩子结点之后，post 被调用
func ForEachNode(n *html.Node, pre, post func(n *html.Node))  {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func StartElement(n *html.Node)  {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth ++
	}
}

func EndElement(n *html.Node)  {
	if n.Type == html.ElementNode {
		depth --
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
