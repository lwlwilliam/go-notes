package main

import "fmt"

// 图 graph 的 key 类型是一个字符串，value 类型 map[string]bool 代表一个字符串集合。
// 从概念上讲，graph 将一个字符串类型的 key 映射到一组相关的字符串集合，它们指向新的 graph 的 key
var graph = make(map[string]map[string]bool)

func addEdge(from, to string)  {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool  {
	return graph[from][to]
}

func main()  {
	addEdge("a", "b")
	fmt.Println(graph)
	fmt.Println(hasEdge("a", "b"))
	graph["c"] = nil
	fmt.Println(graph)
}
