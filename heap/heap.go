package heap

type Heap struct {
	binaryArr []int		// 堆二叉数组
	size int				// 堆规模
}

func NewHeap(cap int) *Heap {
	heap := Heap{
		binaryArr: make([]int, cap),
		size: 0,
	}
	return &heap
}

// O(logN)
func (heap *Heap) Push(v int)  {
	index := heap.size
	heap.size++
	for index > 0 {
		// 父节点编号
		parentIndex := (index-1)/2
		if heap.binaryArr[parentIndex] <= v {
			break
		}
		// 父节点下潜
		heap.binaryArr[index] = heap.binaryArr[parentIndex]
		index = parentIndex
	}
	heap.binaryArr[index] = v
}


// 删除最小(堆顶)元素 1. 将堆尾元素调至堆顶 2. 调整小顶堆
// O(logN)
func (heap *Heap) Pop() int {
	head := heap.binaryArr[0]
	heap.size -= 1
	// 要提到根的数值
	tail := heap.binaryArr[heap.size]
	index := 0
	// 从根开始向下交换
	for index * 2 + 1 < heap.size {
		// 左右孩子编号
		ci, rci := index * 2 + 1, index * 2 + 2
		if rci < heap.size && heap.binaryArr[rci] < heap.binaryArr[ci] {
			ci = rci
		}
		// 没有大小颠倒则退出
		if heap.binaryArr[ci] > tail {
			break
		}
		heap.binaryArr[index] = heap.binaryArr[ci]
		index = ci
	}
	heap.binaryArr[index] = tail
	return head
}