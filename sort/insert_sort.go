package sort

// 插入排序，从位置p开始，进行N-1躺排序, 每次将num[p]插入适当位置
func InsertSort(nums []int) []int {
	for p := 1; p < len(nums); p++ {
		temp := nums[p]
		j := 0
		for j = p; j > 0 && nums[j-1] > temp; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = temp
	}
	return nums
}

