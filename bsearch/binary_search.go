package bsearch

// 二分查找的时间复杂度是 O(logN)，其算法原理十分简单，但是存在很多变形，十个二分九个错

//  最普通的二分查找, 找到在有序数组 nums 中 value 的位置
func BinarySearch(nums []int, value int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		// 使用位运算计算 mid 可以提高速度
		mid := low + ((high - low) >> 1)
		if value < nums[mid] {
			high = mid - 1
		}else if value > nums[mid] {
			low = mid + 1
		}else {
			return mid
		}
	}
	return -1
}

// 在有重复数字的有序数组 nums 中，寻找首次出现 value 的位置
func BinarySearchFirst(nums []int, value int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if value < nums[mid] {
			high = mid - 1
		}else if value > nums[mid] {
			low = mid + 1
		}else {
			if mid == 0 || nums[mid - 1] != value {
				return mid
			}else {
				high = mid - 1
			}
		}
	}
	return -1
}


func BinarySearchLast(nums []int, value int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if value < nums[mid] {
			high = mid - 1
		}else if value > nums[mid] {
			low = mid + 1
		}else {
			if mid == len(nums) - 1 || nums[mid + 1] != value {
				return mid
			}else {
				low = mid + 1
			}
		}
	}
	return -1
}

// 寻找有序数组中第一个大于等于 value 的位置
func BinarySearchFirstBigger(nums []int, value int) int {
	low, high := 0, len(nums) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if value <= nums[mid] {
			if mid == 0 || nums[mid - 1] < value {
				return mid
			}else {
				high = mid - 1
			}
		}else {
			low = mid + 1
		}
	}
	return -1
}

// 寻找有序数组中最后一个小于等于 value 的位置
func BinarySearchLastSmaller(nums []int, value int) int {
	low, high := 0, len(nums) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if value >= nums[mid] {
			if mid == len(nums) - 1 || nums[mid + 1] > value {
				return mid
			}else {
				low = mid + 1
			}
		}else {
			high = mid - 1
		}

	}
	return -1
}