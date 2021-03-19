package list

import "math/rand"

type slNode struct {
	data     int
	forwards []*slNode
	level    int
}

func newSlnode(d int, fl int, l int) *slNode {
	return &slNode{
		data:     d,
		forwards: make([]*slNode, fl),
		level:    l,
	}
}

type skipList struct {
	head      *slNode
	height    int
	maxLevel  int
	skipListP float64
}

func newSkipList(h int) *skipList {
	return &skipList{
		head:      newSlnode(0, h, 0),
		height:    h,
		maxLevel:  1,
		skipListP: 0.5,
	}
}

func (l *skipList) find(val int) *slNode {
	pN := l.head
	for i := l.maxLevel - 1; i >= 0; i-- {
		for pN.forwards[i] != nil && pN.forwards[i].data < val {
			pN = pN.forwards[i]
		}
	}
	if pN.forwards[0] != nil && pN.forwards[0].data == val {
		return pN.forwards[0]
	}
	return nil
}

func (l *skipList) insert(val int) {
	level := l.randomLevel()
	newNode := newSlnode(val, l.height, level)
	update := make([]*slNode, level)
	for i := 0; i < level; i++ {
		update[i] = l.head
	}

	pN := l.head
	for i := level - 1; i >= 0; i-- {
		for pN.forwards[i] != nil && pN.forwards[i].data < val {
			pN = pN.forwards[i]
		}
		update[i] = pN
	}

	for i := 0; i < level; i++ {
		newNode.forwards[i] = update[i].forwards[i]
		update[i].forwards[i] = newNode
	}

	if level > l.maxLevel {
		l.maxLevel = level
	}
}

func (l *skipList) delete(val int) {
	update := make([]*slNode, l.maxLevel)
	pN := l.head
	for i := l.maxLevel - 1; i >= 0; i-- {
		for pN.forwards[i] != nil && pN.forwards[i].data < val {
			pN = pN.forwards[i]
		}
		update[i] = pN
	}
	if pN.forwards[0] != nil && pN.forwards[0].data == val {
		for i := l.maxLevel - 1; i >= 0; i-- {
			if update[i].forwards[i] != nil && update[i].forwards[i].data == val {
				update[i].forwards[i] = update[i].forwards[i].forwards[i]
			}
		}
	}
}

func (l *skipList) randomLevel() int {
	level := 1
	for i := 0; rand.Intn(2) == 1 && i < l.height; i++ {
		level++
	}
	return level
}
