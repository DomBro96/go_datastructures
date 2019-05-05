package heap

// 堆排序
// 1. 将数组变成大顶堆
// 2. 依次将堆顶元素和堆尾元素交换

func HeapSort(array []int)  {
	length := len(array)
	for i := length/2; i >= 0; i-- {
		doBuildHeap(array, i, length)
	}
	for i := length-1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		doBuildHeap(array, 0, i)
	}
}



func doBuildHeap(array []int, pi, length int)  {
	index := pi
	parent := array[pi]
	for 2*index+1 < length {
		// 左右孩子编号
		ci, rci := 2*index+1, 2*index+2
		if rci < length && array[ci] < array[rci] {
			ci = rci
		}
		if parent >= array[ci] {
			break
		}
		if parent < array[ci] {
			array[index] = array[ci]
		}
		index = ci
	}
	array[index] = parent
}