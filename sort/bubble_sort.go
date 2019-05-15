package sort

// 冒泡排序
func BubbleSort(nums []int) []int  {
	for i := 1; i < len(nums); i++{
		for j := 0; j < len(nums)-i; j++ {
			if j+1 < len(nums) && nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
