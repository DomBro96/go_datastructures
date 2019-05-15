package sort

func MergeSort(nums []int) []int  {
	tempNums := make([]int, len(nums))
	mergeSort(0, len(nums)-1, nums, tempNums)
	return nums
}

func mergeSort(left, right int, nums, tempNums []int)  {
	if left < right {// 对前一半数组和后一半数组进行排序，然后合并两个有序数组
		center := (left+right)/2
		mergeSort(left, center, nums, tempNums)
		mergeSort(center+1, right, nums, tempNums)
		merge(left, center, center+1, right, nums, tempNums)
	}
}
// 合并两个有序数组
func merge(leftPos, leftEnd, rightPos, rightEnd int, nums, tempNums []int) {
	tempPos, leftOri := leftPos, leftPos
	for leftPos <= leftEnd && rightPos <= rightEnd {
		if nums[leftPos] <= nums[rightPos] {
			tempNums[tempPos] = nums[leftPos]
			leftPos++
			tempPos++
		}else {
			tempNums[tempPos] = nums[rightPos]
			rightPos++
			tempPos++
		}
	}
	for leftPos <= leftEnd {
		tempNums[tempPos] = nums[leftPos]
		leftPos++
		tempPos++
	}
	for rightPos <= leftEnd {
		tempNums[tempPos] = nums[rightPos]
		rightPos++
		tempPos++
	}
	for i := tempPos-1; i >= leftOri; i-- {
		nums[i] = tempNums[i]
	}
}
