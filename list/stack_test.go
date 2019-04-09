package list

import (
	"fmt"
	"testing"
)

func TestSliceStack_IsEmpty(t *testing.T) {
	stack := &SliceStack{
		elements: make([]int, 0),
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	for j := 0; j < 5; j++ {
		i, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(i)
	}

}

func TestLinkedStack_IsEmpty(t *testing.T) {
	stack := NewLinkedStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	for j := 0; j < 5; j++ {
		i, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(i)
	}
}

