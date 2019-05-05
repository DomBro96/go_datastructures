package heap

import (
	"fmt"
	"testing"
)

func TestHeap_Pop(t *testing.T) {
	// 测试优先级队列
	heap := NewHeap(5)
	heap.Push(5)
	heap.Push(4)
	heap.Push(3)
	heap.Push(2)
	heap.Push(1)
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
}


func TestHeap_Pop2(t *testing.T) {
	// 测试优先级队列
	heap := NewHeap(5)
	heap.Push(5)
	heap.Push(5)
	heap.Push(5)
	heap.Push(2)
	heap.Push(1)
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Pop())
}