package dp

import "fmt"

var (
	w = []int{2, 1, 3, 2}
	v = []int{3, 2, 4, 2}
	N = 4
)

// 背包问题递归求解

// 从第i个物品开始，总量小于j的部分
func BagRec(i, j int) int  {
	res := 0
	if i == N {
		res = 0
	} else if w[i] > j {  // 超重查看下一个物品
		res = BagRec(i+1, j)
	}else {
		res = max(BagRec(i+1, j), BagRec(i+1, j - w[i])+v[i])		// 挑选和不挑选都尝试一下
	}
	return res
}

// 背包问题动态规划解法 逆向求解
func BagDp(weight int) int {
	// dp[i][j]  从第 i 个物品开始选取重量为 j 的最大价值
	dp := make([][]int, 0)
	for i := 0; i <= N; i++ {
		dp = append(dp, make([]int, weight+1))
	}
	for i := N-1; i >= 0; i--{
		for j := 0; j <= weight; j++ {
			if w[i] > j {
				dp[i][j] = dp[i+1][j]
			}else {
				dp[i][j] = max(dp[i+1][j], dp[i+1][j-w[i]]+v[i])
			}
		}
	}
	fmt.Println(dp)
	return dp[0][weight]
}

// 背包问题 正向动态规划求解
func BagDp2(weight int) int {
	//dp[i+1][j] 前i个物品重量不超过j的最大价值
	dp := make([][]int, N+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, weight)
	}
	for i := 0; i <= N; i++ {
		for j := 0; j <= weight; j++ {
			if j < w[i] {
				dp[i+1][j] = dp[i][j]
			}else {
				dp[i+1][j] = max(dp[i][j], dp[i][j-w[i]]+v[i])
			}
		}
	}
	return dp[N][weight]
}

func max(a, b int) int {
	if a >= b {
		return a
	}else {
		return b
	}
}


// 背包问题，每个物品可以任意选择个数
func BagDp3(weight int) int {
	//dp[i+1][j] 前i个物品重量不超过j的最大价值
	dp := make([][]int, N+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, weight)
	}
	for i := 0; i <= N; i++ {
		for j := 0; j <= weight; j++ {
			for k := 0; k * w[i] <= j; k++ {
				dp[i+1][j] = max(dp[i+1][j], dp[i][j - k*w[i]] + k*v[i])
			}
		}
	}
	return dp[N][weight]
}