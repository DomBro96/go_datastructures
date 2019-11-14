package graph

/**
	部分和问题，给定整数 a1...an, 判断能否从整数中选出若干个数，使得和为 k。 这是一道深度优先搜索问题，
	深度优先搜索存在着状态的转移，当满足某个状态时，继续下一步，否则回退到上一步
 */



func PartialSummation(nums []int, sum int) bool  {
	return Dfs(0, nums, sum, 0)
}

func Dfs(index int, nums []int, sum int, partSum int) bool {
	if index == len(nums) - 1 {
		return partSum == sum
	}
	// 不带上该数字
	if Dfs(index + 1, nums, sum, partSum) {
		return true
	}
	// 带上该数字
	if Dfs(index + 1, nums, sum, partSum + nums[index]) {
		return true
	}
	return false
}