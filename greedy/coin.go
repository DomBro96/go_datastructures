package greedy


// 硬币问题: 给定系列面值硬币 {1，5，10，20，100，500} 对应个数为 {C1，C5，C10，C20，C100，C200}
// 用这些硬币购买价格为 A 的商品，求最所需最少硬币数。
// 贪心算法是遵循某种规律，一直不断向前求解当前最优策略的算法，硬币问题的规律是选取面值最大的硬币。

var (
	// 硬币
	coins = []int{1, 5, 10, 50, 100, 500}
)

func CoinSolve(num []int, price int) int {
	res := 0
	for i := len(coins)-1; i >= 0; i-- {
		// 从最大面值开始
		count := min(price/coins[i], num[i])
		price -= count*coins[i]
		res += count
	}
	return res
}

func min(a, b int) int  {
	if a <= b {
		return a
	}else {
		return b
	}
}