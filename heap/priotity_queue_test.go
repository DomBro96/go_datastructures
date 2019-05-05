package heap

import (
	"fmt"
	"testing"
)

func TestPriorityQueue_Pop(t *testing.T) {
	queue := NewPriorityQueue(5)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())

}