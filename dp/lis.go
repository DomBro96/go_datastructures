package dp

/**
	最长上升子序列问题，上升子序列是指任意的 i < j 都满足 ai < aj 的序列
	dp[i] 表示以 i 结尾的子序列最长的上升序列
 */

func Lis(nums []int) int {
	res := 0
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}
