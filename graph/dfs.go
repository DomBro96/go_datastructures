package graph

/**
部分和问题，给定整数 a1...an, 判断能否从整数中选出若干个数，使得和为 k。 这是一道深度优先搜索问题，
深度优先搜索存在着状态的转移，当满足某个状态时，继续下一步，否则回退到上一步
*/

func PartialSummation(nums []int, sum int) bool {
	return Dfs(0, nums, sum, 0)
}

func Dfs(index int, nums []int, sum int, partSum int) bool {
	if index == len(nums)-1 {
		return partSum == sum
	}
	// 不带上该数字
	if Dfs(index+1, nums, sum, partSum) {
		return true
	}
	// 带上该数字
	if Dfs(index+1, nums, sum, partSum+nums[index]) {
		return true
	}
	return false
}

func FullPermutation(words []string) [][]string {

	stack := make([]string, 0, len(words))
	ps := make([][]string, 0)

	permutation(words, stack, ps)

	return ps

}

func permutation(words []string, stack []string, ps [][]string) {
	if len(words) <= 0 {

		cp := make([]string, 0, len(stack))
		copy(cp, stack)
		ps = append(ps, cp)

	} else {

		for i := 0; i < len(words); i++ {

			ws := make([]string, 0, len(words)-1)
			ws = append(ws, words[0:i]...)
			if i+1 < len(words) {
				ws = append(ws, words[i+1:]...)
			}

			stack = append(stack, words[i])
			permutation(ws, stack, ps)
			stack = stack[:len(stack)-1]
		}
	}
}
