// 数据结构之双向链表
// 什么是双向链表？和单链表比较，双向链表的节点不但知道自己的下线，还知道自己的上线。每个节点除了一个指向后面节点的指针外，还有一个指向前面
// 节点的指针（链表头和尾除外）。链表头只有指向后面节点的指针，链表尾只有指向前面节点的指针。
package dataStructure

type DNode struct {
	Item	Object
	prev	*DNode
	next	*DNode
}

type DList struct {
	size	uint64
	head	*DNode
	tail	*DNode
}

// 初始化
func (dList *DList) Init() {
	*dList = DList{
		size: 0,
		head: nil,
		tail: nil,
	}
}

// 新增数据
func (dList *DList) Append(dNode *DNode) {
	if dList.size == 0 {
		dNode.prev = nil
		dNode.next = nil
		dList.head = dNode
		dList.tail = dNode
	} else {
		dNode.prev = dList.tail
		dNode.next = nil
		dList.tail.next = dNode
		dList.tail = dNode
	}

	dList.size ++
}
