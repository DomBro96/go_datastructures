package list

import (
	"fmt"
	"testing"
)

func TestSliceQueue_IsEmpty(t *testing.T) {
	queue := NewSliceQueue()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)

	for j := 0; j < 5; j++ {
		i, err := queue.DeQueue()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(i)
	}
}

func TestLinkedQueue_IsEmpty(t *testing.T) {
	queue := NewLinkedQueue()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)
}

