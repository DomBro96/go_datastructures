package heap

// 堆 (小根堆数据结构，数据类型试用简单的 int，优先级队列)
type Heap struct {
	Capacity int
	Size 	 int
	Elements []int
}

func NewHeap(maxCapacity int) *Heap  {
	if maxCapacity < 1 {
		return nil
	}
	heap := &Heap{
		Capacity: maxCapacity,
		Size: 0,
		Elements: make([]int, maxCapacity),
	}
	return heap
}

// 使用 上滤 策略实现插入, 新建空穴， 空穴上潜
func (heap *Heap)Insert(val int)  {
	var i int
	if heap.IsFull() {
		return
	}
	heap.Size += 1
	for i = heap.Size; heap.Elements[i / 2] > val; i /= 2 {
		heap.Elements[i] = heap.Elements[i / 2]
	}
	heap.Elements[i] = val
}

// 删除根节点 下滤 策略实现删除， 新建空穴， 空穴下潜
func (heap *Heap)DeleteMin() (int, bool) {
	var i, child int
	var minElement, lastElement int
	if heap.IsEmpty() {
		return -1, false
	}
	for i = 0; i * 2 < heap.Size; i = child {
		child = i * 2
		if child != heap.Size && heap.Elements[child + 1] < heap.Elements[child] {
			child++
		}
		// 把较小的儿子置入空穴， 空穴向下移动一层
		if lastElement > heap.Elements[child] {
			heap.Elements[i] = heap.Elements[child]
		}else {
			break
		}
	}
	heap.Elements[i] = lastElement
	return minElement, true
}

// 直接返回根节点 时间复杂度 O(1)
func (heap *Heap)FindMin() int {
	return heap.Elements[0]
}

// 堆是否为空
func (heap *Heap)IsEmpty() bool {
	return heap.Size == 0
}

// 堆是否已满
func (heap *Heap)IsFull() bool  {
	return heap.Size == heap.Capacity
}


