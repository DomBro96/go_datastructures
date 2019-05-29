package sort

// 基于快速排序，寻找数组中第K个大的数字
// f(0...n) = f(0...p) + f(p+1...) p+1 != K
// p+1 = K  return nums[p]
func FindKthNum(nums []int, K int) int {
	if  K > len(nums) || K < 1 {
		return -1
	}
	return doFindKthNum(0, len(nums)-1, K, nums)
}

func doFindKthNum(left, right, K int, nums []int) int {
	// 枢纽元素
 	if right > left {
		pivot := nums[right]
		i, j := left, right - 1
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
		nums[i], nums[right] = nums[right], nums[i]
		p := i+1
		if p == K {
			return nums[i]
		}else if p < K {
			return doFindKthNum(i+1, right, K, nums)
		}else {
			return doFindKthNum(left, i-1, K, nums)
		}
	}else {
		if left + 1 == K {
			return nums[left]
		}else {
			return -1
		}
	}
}
