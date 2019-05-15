package sort

// 堆排序，构造大顶堆，交换堆顶元素和堆尾元素

func HeapSort(nums []int) []int  {
	length := len(nums)
	// 从第N/2个节点开始下潜，保证每一个父节点大于其儿子节点, 构建大顶堆
	for i := length/2; i >= 0; i-- {
		heapHelper(nums, i, length)
	}
	// 每次交换堆顶和堆尾的元素， 重新构建大顶堆
	for i := 1; i < length; i++ {
		nums[0], nums[length-i] = nums[length-i], nums[0]
		heapHelper(nums, 0, length-i)
	}
	return nums
}

// 将父节点下潜至合适位置
// pi 父节点编号
func heapHelper(nums []int, pi, length int) {
	// 父节点的值
	parent := nums[pi]
	for 2*pi + 1 < length {
		ci, rci := 2*pi+1, 2*pi+2
		if rci < length && nums[rci] > nums[ci] {
			ci = rci
		}
		if parent >= nums[ci] {
			break
		}else {
			// 儿子节点上浅
			nums[pi] = nums[ci]
			pi = ci
		}
	}
	nums[pi] = parent
}

