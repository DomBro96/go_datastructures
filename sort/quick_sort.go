package sort

// 快速排序
// 三数中值法选出一个数字m， 比该数字小的，放置在该数字的左边，比该数字大的放置在该数字的右边
func QuickSort(nums []int) []int {
	quickSort(0, len(nums)-1, nums)
	return nums
}

func quickSort(left, right int, nums []int)  {
	// 快速排序适合较大规模的数组排序
	if right - left > 20 {
		// 枢纽元素
		pivot := median(left, right, nums)
		nums[(left+right)/2], nums[right-1] = nums[right-1], nums[(left+right)/2]
		i := left + 1
		j := right - 2
		for {
			for nums[i] < pivot {
				i++
			}
			for nums[j] > pivot {
				j--
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			}else {
				break
			}
		}
		nums[i], nums[right-1] = nums[right-1], nums[i]
		quickSort(left, i-1, nums)
		quickSort(i+1, right, nums)
	}else {
		InsertSort(nums)
	}

}

// 三数中值法
func median(left, right int, nums []int) int {
	center := (left+right)/2
	if nums[left] > nums[right] {
		nums[left], nums[right] = nums[right], nums[left]
	}
	if nums[left] > nums[center] {
		nums[left], nums[center] = nums[center], nums[left]
	}
	if nums[center] > nums[right] {
		nums[center], nums[right] = nums[right], nums[center]
	}
	return nums[center]
}
