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

func TestQueueStack_Push(t *testing.T) {
	qStack := NewQueueStack()
	qStack.Push(1)
	qStack.Push(2)
	qStack.Push(3)
	qStack.Push(4)
	qStack.Push(5)

	for !qStack.IsEmpty() {
		n, err := qStack.Pop()
		t.Logf("pop: %d, err: %v", n, err)
	}
}

func TestQueueStack_Top(t *testing.T) {
	qStack := NewQueueStack()
	qStack.Push(1)
	top, err := qStack.Top()
	t.Logf("top: %d, err: %v", top, err)
	qStack.Push(2)
	top, err = qStack.Top()
	t.Logf("top: %d, err: %v", top, err)
	qStack.Push(3)
	top, err = qStack.Top()
	t.Logf("top: %d, err: %v", top, err)
	qStack.Push(4)
	top, err = qStack.Top()
	t.Logf("top: %d, err: %v", top, err)
	qStack.Push(5)
	top, err = qStack.Top()
	t.Logf("top: %d, err: %v", top, err)

}
