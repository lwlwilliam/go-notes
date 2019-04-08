// 使用 slice 来模拟 stack
package sliceStack

import (
	"errors"
	"log"
)

type Stack []int

func (s *Stack) Push(v int)  {
	*s = append(*s, v)
}

// 弹出最后一个元素
func (s *Stack) Pop() error {
	if len(*s) < 1 {
		return errors.New("the Stack is empty")
	}
	*s = (*s)[:len(*s) - 1]

	return nil
}

// 删除指定元素
func (s *Stack) Remove(i int) (Stack, error) {
	log.Printf("index: %d, len: %d\n", i, len(*s))
	if i < 0 || len(*s) <= i {
		log.Println("invalid index")
		return nil, errors.New("invalid index")
	}

	s2 := make(Stack, len(*s) - 1)
	copy((*s)[i:], (*s)[i+1:])
	copy(s2, *s)

	return s2, nil
}
