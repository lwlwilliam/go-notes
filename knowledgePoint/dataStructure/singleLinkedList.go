// 数据结构之单链表
package dataStructure

// 数据
type Object interface{}

// 节点结构，每个节点都有指向下一个节点的指针 *Next，还有附带的数据 Object
type Node struct {
	Item	Object
	Next	*Node
}

// 链表结构
type List struct {
	size	uint64	// 节点数量
	head	*Node	// 链表头
	tail	*Node	// 链表尾
}

// 初始化
func (list *List) Init() {
	*list = List{
		size: 0,
		head: nil,
		tail: nil,
	}
}

// 添加元素
func (list *List) Append(node *Node) bool {
	if node == nil {
		return false
	}

	(*node).Next = nil

	// 首次添加指定链表头
	if (*list).size == 0 {
		(*list).head = node
	// 注意，先为添加前的链表尾部节点指定下一个节点的指针
	} else {
		(*list).tail.Next = node
	}

	// 注意，先为添加前的链表尾部节点指定下一个节点的指针才能把尾部指向新添加的节点
	(*list).tail = node
	(*list).size ++

	return true
}

// 插入元素，在 list 的索引为 i 的位置插入 node 节点
func (list *List) Insert(i uint, node *Node) bool {
	// node 不能为空、索引不能大于 list 长度、list 不能为空
	if node == nil || i > uint((*list).size) || (*list).size == 0 {
		return false
	}

	// 首部插入
	if i == 0 {
		// node 节点的 Next 指针指原先的 head，把 head 指向新 node 节点
		(*node).Next = (*list).head
		(*list).head = node
	} else {
		// 从首节点开始遍历，找出第 i - 1 个节点
		preNode := (*list).head
		for j := uint(1); j < i; j ++ {
			preNode = (*preNode).Next
		}

		// 把 node 节点的 Next 指针指向原先 i 位置的节点
		(*node).Next = (*preNode).Next
		// i - 1 位置的 Next 指针指向新的 node 节点
		(*preNode).Next = preNode
	}

	(*list).size ++

	return true
}

// 删除元素，其实 node 这个参数没什么用啊
func (list *List) Remove(i uint, node *Node) bool {
	if i >= uint((*list).size) {
		return false
	}

	if i == 0 {
		(*list).head = (*list).head.Next
		if (*list).size == 1 {
			(*list).tail = nil
		}
	} else {
		preNode := (*list).head
		// 找出第 i-1 个元素
		for j := uint(1); j < i; j ++ {
			preNode = (*preNode).Next
		}

		// 把第 i 个节点的前一个节点的 Next 指针直接指向第 i+1 个元素
		(*preNode).Next = (*preNode).Next.Next

		// 删除尾部的情况
		if i == (uint((*list).size) - 1) {
			(*list).tail = preNode
		}
	}

	(*list).size --

	return true
}

// 获取索引 i 位置的元素
func (list *List) Get(i uint) *Node {
	if i >= uint((*list).size) {
		return nil
	}

	node := (*list).head
	for j := uint(0); j < i; j ++ {
		node = (*node).Next
	}

	return node
}

// 链表长度
func (list *List) Size() uint64 {
	return list.size
}

// 清空链表
func (list *List) RemoveAll() {
	*list = List{
		size: 0,
		head: nil,
		tail: nil,
	};
}
