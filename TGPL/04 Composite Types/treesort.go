package main

import (
	"fmt"
)

type tree struct {
	value int
	left, right *tree
}

// Sort sorts values in place.
// 对 slice 中的元素进行排序
// 先遍历 slice，并把元素添加到 二叉树中
// 调用 appendValues 函数，把二叉树中的元素进行排序并添加到 slice 中
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	var s = []int{8, 1, 18, 23, 6, 12}
	Sort(s)
	fmt.Println(s)
}
