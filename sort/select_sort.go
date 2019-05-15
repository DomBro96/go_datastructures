package sort

func SelectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		chosen := i
		for j := i+1; j < len(nums); j++ {
			if nums[j] < nums[chosen] {
				chosen = j
			}
		}
		if chosen != i {
			nums[chosen], nums[i] = nums[i], nums[chosen]
		}
	}
	return nums
}
