package list

import (
	"fmt"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	n :=  &LinkedNode{
		val: 1,
		next: nil,
	}
	n1 := &LinkedNode{
		val: 2,
		next: nil,
	}
	n2 := &LinkedNode{
		val: 3,
		next: nil,
	}
	n3 := &LinkedNode{
		val: 4,
		next: nil,
	}
	n4 := &LinkedNode{
		val: 5,
		next: nil,
	}
	n.next = n1
	n1.next = n2
	n2.next = n3
	n3.next = n4
	head := ReverseLinkedList(n)
	pNode := head
	for pNode != nil {
		fmt.Println(pNode.val)
		pNode = pNode.next
	}

}
