package list

import "errors"

type Stack interface {
	Push(i int)
	Pop() (int, error)
	Top() int
	IsEmpty() bool
}


type SliceStack struct {
	elements []int
}

func NewSliceStack(c int) *SliceStack {
	return &SliceStack{
		elements: make([]int, 0, c),
	}
}

func (s *SliceStack) Push(i int)  {
	s.elements = append(s.elements, i)
}

func (s *SliceStack) Pop() (int, error) {
	if !s.IsEmpty() {
		top := s.elements[len(s.elements) - 1]
		s.elements = s.elements[:len(s.elements) - 1]
		return top, nil
	}else {
		return -1, errors.New("stack is empty. ")
	}
}

func (s *SliceStack) Top() (int, error)  {
	if !s.IsEmpty() {
		top := s.elements[len(s.elements) - 1]
		return top, nil
	}else {
		return -1, errors.New("stack is empty. ")
	}
}

func (s *SliceStack) IsEmpty() bool  {
	return len(s.elements) == 0
}

type Node struct {
	val int
	next *Node
}

type LinkedStack struct {
	head *Node
	size int
}

func NewLinkedStack() *LinkedStack  {
	return &LinkedStack{
		head: &Node{
			val: 0,
			next: nil,
		},
		size: 0,
	}
}

func (s *LinkedStack) Push(i int) {
	newNode := &Node{
		val: i,
		next: nil,
	}
	if s.IsEmpty() {
		s.head.next = newNode
	}else {
		curNode := s.head
		for i := 0; i < s.size; i++ {
			curNode = curNode.next
		}
		curNode.next = newNode
	}
	s.size += 1
}
func (s *LinkedStack) Pop() (int, error) {
	if !s.IsEmpty() {
		curNode := s.head
		nextNode := curNode.next
		for nextNode.next != nil {
			curNode = nextNode
			nextNode = nextNode.next
		}
		popNode := curNode.next
		curNode.next = nil
		s.size -= 1
		return popNode.val, nil
	}else {
		return -1, errors.New("stack is empty. ")
	}
}
func (s *LinkedStack) Top() (int, error) {
	if !s.IsEmpty() {
		curNode := s.head
		nextNode := curNode.next
		for nextNode.next != nil {
			curNode = nextNode
			nextNode = nextNode.next
		}
		topNode := curNode.next
		return topNode.val, nil
	}else {
		return -1, errors.New("stack is empty. ")
	}
}

func (s *LinkedStack) IsEmpty() bool {
	return s.size == 0
}