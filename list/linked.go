package list

// 链表结构相关训练


type (
	// 单链表节点
	LinkedNode struct {
		val  int
		next *LinkedNode
	}

	// 双向链表节点
	DoublyLinkedNode struct {
		val  int
		pre  *LinkedNode
		next *LinkedNode
	}
)

// 反转单链表
func ReverseLinkedList(head *LinkedNode) *LinkedNode {
	// 当前节点
	pNode := head
	// 前驱节点
	preNode := &LinkedNode{}
	for pNode.next != nil {
		nextNode := pNode.next
		pNode.next = preNode
		preNode = pNode
		pNode = nextNode
	}
	pNode.next = preNode
	return pNode
}



