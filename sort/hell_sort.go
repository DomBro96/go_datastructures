package sort

// 希尔排序， 每趟排序保证在增量范围内是有序的
func HellSort(nums []int) []int {
	// 增量
	j := 0
	for gap := len(nums)/2; gap > 0; gap/=2{
		for i := gap; i < len(nums); i++ {
			temp := nums[i]
			for j = i; j >= gap && nums[j-gap] > temp; j-=gap {
				nums[j] = nums[j-gap]
			}
			nums[j] = temp
		}
	}
	return nums
}
