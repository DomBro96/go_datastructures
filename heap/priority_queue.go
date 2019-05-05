package heap

// 大顶堆实现的优先级队列
type PriorityQueue struct {
	binaryArr []int
	size int
}

func NewPriorityQueue(cap int) *PriorityQueue {
	queue := &PriorityQueue{
		binaryArr: make([]int, cap),
		size: 0,
	}
	return queue
}

// 向优先级队列插入元素
func (queue *PriorityQueue) Push(v int)  {
	// 插入编号
	index := queue.size
	queue.size += 1
	for index > 0 {
		// 父节点编号
		parentIndex := (index-1)/2
		if queue.binaryArr[parentIndex] >= v {
			break
		}
		// 父节点下潜
		queue.binaryArr[index] = queue.binaryArr[parentIndex]
		index = parentIndex
	}
	queue.binaryArr[index] = v
}

// 删除最大(优先级最高)元素 1.将堆尾元素调至堆顶  2.调整大顶堆
func (queue *PriorityQueue) Pop() int  {
	 head := queue.binaryArr[0]
	 queue.size -= 1
	 // 堆尾元素
	 tail := queue.binaryArr[queue.size]
	 index := 0
	 for index*2+1 < queue.size {
	 	// 左右孩子编号
	 	ci, rci := index*2+1, index*2+2
	 	// 找到更大的那一个编号
	 	if rci < queue.size && queue.binaryArr[rci] > queue.binaryArr[ci] {
	 		ci = rci
		}
	 	if tail >= queue.binaryArr[ci] {
	 		break
		}
	 	// 儿子节点上浅
	 	queue.binaryArr[index] = queue.binaryArr[ci]
	 	index = ci
	 }
	 queue.binaryArr[index] = tail

	 return head
}
