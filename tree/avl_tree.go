package tree

/**
	AVL 树是自平衡的树， 每隔节点的左右子树高度差不超过一
 */

const ALLOW_BLANCE = 1

type AvlNode struct {
	left  *AvlNode
	right *AvlNode
	val int
	height int
}

type AvlTree = AvlNode



func NewAvlTree(rVal int) *AvlTree {
	return &AvlNode{
		left: nil,
		right: nil,
		val: rVal,
		height: 0,
	}
}

func (t *AvlTree) Insert(val int) *AvlTree {
	if isExist, _ := t.Find(val); isExist {
		return t
	}
	return insertAvlNode(val, t)
}

func insertAvlNode(val int, node *AvlNode) *AvlNode {
	if node == nil {
		node = &AvlNode{
			val: val,
			left: nil,
			right: nil,
			height: 0,
		}
	}else {
		if val < node.val {
			node.left = insertAvlNode(val, node.left)
		}else if val > node.val {
			node.right = insertAvlNode(val, node.right)
		}
	}
	// 插入后对树进行平衡
	return balance(node)
}

func (t *AvlTree) Del(val int) bool  {
	isExist, _ := t.Find(val)
	if ! isExist {
		return false
	}
	delAvlNode(val, t)
	return true

}

// 1. 删除节点是叶子节点，直接删除
// 2. 删除节点具有单个子节点，其子节点替代它的位置
// 3. 删除节点具有两个子节点，使用其右子树中最小的节点代替
// 4. 每次删除要进行一次平衡
func delAvlNode(val int, node *AvlNode) *AvlNode {
	if val < node.val {
		node.left = delAvlNode(val, node.left)
	}else if val > node.val {
		node.right = delAvlNode(val, node.right)
	}else if node.left != nil && node.right != nil {
		rValue, _ := node.right.FindMin()
		node.val = rValue
		node.right = delAvlNode(rValue, node)
	}else {
		if node.left != nil {
			node = node.left
		}else {
			node = node.right
		}
	}
	return balance(node)
}

func (t *AvlNode) FindMin() (int, *AvlNode) {
	curNode := t
	for curNode.left != nil {
		curNode = curNode.left
	}
	return curNode.val, curNode
}

// 平衡方法, 参数为可能的非平衡节点
func balance(node *AvlNode) *AvlNode  {
	if node == nil {
		return node
	}

	if high(node.left) - high(node.right) > ALLOW_BLANCE {
		if high(node.left.left) >= high(node.left.right) {
			// 左-左
			singleRotateWithLeft(node)
		}else {
			// 左-右
			doubleRotateWithLeft(node)
		}
	}else if high(node.right) - high(node.left) > ALLOW_BLANCE {
		if high(node.right.right) >= high(node.right.left) {
			// 右-
			singleRotateWithRight(node)
		}else {
			// 右-左
			doubleRotateWithRight(node)
		}
	}
	node.height = max(high(node.left), high(node.right)) + 1
	return node
}

func (t *AvlTree)Find(val int) (bool, *AvlNode) {
	return findAvlNode(val, t)
}


func findAvlNode(val int, node *AvlNode) (bool, *AvlNode)  {
	if node == nil {
		return false, nil
	}else {
		if val < node.val {
			return findAvlNode(val, node.left)
		}else if val > node.val {
			return findAvlNode(val, node.right)
		}else {
			return true, node
		}
	}
}

func high(node *AvlNode) int {
	if node == nil {
		return -1
	}else {
		return node.height
	}
}

// 左左单旋转 向不平衡节点左孩子的左子树插入新节点
// 将 不平衡节点 k2 左子节点 k1 向上拖拽， 使 k2 作为 k1 的右子节点， k1 的 右子节点  作为 k2 的左子节点。
func singleRotateWithLeft(k2 *AvlNode) *AvlNode{
	k1 := k2.left
	k2.left = k1.right
	k1.right = k2
	k1.height = max(high(k1.left), k2.height) + 1
	k2.height = max(high(k2.left), high(k2.right)) + 1
	return k1											// new root
}

// 右右单旋转 向不平衡节点右孩子的右子树插入新节点
// 将不平衡节点 k2 的右子节点 k1 向上拖拽， 使 k2 作为 K1 的左子节点， k1 的 左子节点 作为 k2 的右子节点。
func singleRotateWithRight(k2 *AvlNode) *AvlNode  {
	k1 := k2.right
	k2.right = k1.left
	k1.left = k2
	k1.height = max(k2.height, high(k1.right)) + 1
	k2.height = max(high(k2.left), high(k2.right)) + 1
	return k1										// new root
}

// 左右双旋转 向不平衡节点的左孩子的右子树插入新节点
// k2 作为不平衡节点， 首先将 k2 的 左儿子 进行 右右单旋转， 然后对 k2 进行 左左 单旋转
func doubleRotateWithLeft(k2 *AvlNode) *AvlNode {
	singleRotateWithRight(k2.left)
	return singleRotateWithLeft(k2)
}


// 右左双旋转 向不平衡节点的右孩子的左子树插入新节点
// k2 作为不平衡节点， 首先将 k2 的 右儿子 进行左左但旋转， 然后对 k2 进行 右右 单旋转
func doubleRotateWithRight(k2 *AvlNode) *AvlNode {
	singleRotateWithLeft(k2.left)
	return singleRotateWithRight(k2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
