package links

import (
	"golang.org/x/net/html"
	"fmt"
)

// 每遇到一个 HTML 元素标签，就将其入栈，并输出
func Outline(stack []string, n *html.Node)  {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)	// push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Outline(stack, c)
	}
}
