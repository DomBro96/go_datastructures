package list

import "errors"

type Queue interface {
	EnQueue(i int)
	DeQueue() (int, error)
	Front() (int, error)
	IsEmpty() bool
}

type SliceQueue struct {
	elements []int
}

func NewSliceQueue() *SliceQueue {
	return &SliceQueue{
		elements: make([]int, 0),
	}
}

func (q *SliceQueue) EnQueue(i int)  {
	q.elements = append(q.elements, i)
}

func (q *SliceQueue) DeQueue() (int, error)  {
	if ! q.IsEmpty() {
		front := q.elements[0]
		q.elements = q.elements[1:]
		return front, nil
	} else {
		return -1, errors.New("queue is empty. ")
	}
}

func (q *SliceQueue) Front() (int, error)  {
	if !q.IsEmpty() {
		front := q.elements[0]
		return front, nil
	} else {
		return -1, errors.New("queue is empty. ")
	}
}

func (q *SliceQueue) IsEmpty() bool {
	return len(q.elements) == 0
}

type LinkedQueue struct {
	head *Node
	size int
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{
		head: &Node{
			val: -1,
			next: nil,
		},
		size: 0,
	}
}

func (q *LinkedQueue) EnQueue(i int)  {
	newNode := &Node{
		val: i,
		next: nil,
	}
	if q.IsEmpty() {
		q.head.next = newNode
	}else {
		curNode := q.head
		for i := 0; i < q.size; i++ {
			curNode = curNode.next
		}
		curNode.next = newNode
	}
	q.size += 1
}

func (q *LinkedQueue) DeQueue() (int, error)  {
	if q.IsEmpty() {
		return -1, errors.New("queue is empty. ")
	}else {
		front := q.head.next
		nextNode := front.next
		q.head.next = nextNode
		q.size -= 1
		return front.val, nil
	}
}

func (q *LinkedQueue) Front() (int, error)   {
	if !q.IsEmpty() {
		return -1, errors.New("queue is empty. ")
	}else {
		front := q.head.next
		return front.val, nil
	}
}

func (q *LinkedQueue) IsEmpty() bool  {
	return q.size == 0
}

