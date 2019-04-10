package toposort

import (
	"log"
	"sort"
)

func TopoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				log.Println(item)
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	// 获取所有有前置课程的课程
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	visitAll(keys)
	return order
}
