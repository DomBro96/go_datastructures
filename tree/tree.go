package tree

import (
	"fmt"
)

type BstNode struct {
	left  *BstNode
	right *BstNode
	val int
}

type Bst = BstNode 	// 二叉搜索树别名

func NewBst(rVal int) *Bst{
	return &BstNode{
		val: rVal,
		left: nil,
		right: nil,
	}
}

func (t *Bst)Insert(val int) *Bst {
	if isExist, _ := t.Find(val); isExist {
		return t
	}
	return insert(val, t)
}

func insert(val int, node *BstNode) *BstNode {
	if node == nil {
		node = &BstNode{
			val: val,
			left: nil,
			right: nil,
		}
	}else {
		if val < node.val {
			node.left = insert(val, node.left)
		}else if val > node.val {
			node.right = insert(val, node.right)
		}
	}
	return node
}

func (t *Bst)Find(val int) (bool, *BstNode) {
	return find(val, t)
}

func find(val int, node *BstNode) (bool, *BstNode) {
	if node == nil {
		return false, nil
	}else {
		if val < node.val {
			return find(val, node.left)
		}else if val > node.val {
			return find(val, node.right)
		}else {
			return true, node
		}
	}
}




func (t *Bst) Del(val int) bool {
	isExist, _ := t.Find(val)
	if !isExist {
		return false
	}
	del(val, t)
	return true
}

// 1. 删除节点是叶子节点，直接删除
// 2. 删除节点具有单个子节点，其子节点替代它的位置
// 3. 删除节点具有两个子节点，使用其右子树中最小的节点代替
func del(val int, node *BstNode) *BstNode  {
	if val < node.val {
		node.left = del(val, node.left)
	}else if val > node.val {
		node.right = del(val, node.right)
	}else if node.left != nil && node.right != nil {
		rValue,_ := node.right.FindMin()
		node.val = rValue
		node.right = del(rValue, node.right)
	}else {
		if node.left != nil {
			node = node.left
		}else {
			node = node.right
		}
	}
	return node
}

func (t *Bst)FindMin() (int, *BstNode) {
	curNode := t
	for curNode.left != nil {
		curNode = curNode.left
	}
	return curNode.val, curNode
}


func (t *Bst)FindMax() (int, *BstNode) {
	curNode := t
	for curNode.right != nil {
		curNode = curNode.right
	}
	return curNode.val, curNode
}

func (t *Bst) InOrderDisplay()  {
	if t != nil {
		t.left.InOrderDisplay()
		fmt.Println(t.val)
		t.right.InOrderDisplay()
	}
}

func Insert(node *BstNode, val int) *BstNode {
	if node == nil {
		node.val = val
		node.left = nil
		node.right = nil
	}else {
		if val < node.val {
			node.left = Insert(node.left, val)
		}else if val > node.val {
			node.right = Insert(node.right, val)
		}
	}
	return node
}


