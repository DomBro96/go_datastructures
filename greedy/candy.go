package greedy

import "math"

/**
	分糖果问题： 有 n 个孩子， m 个糖果，n > m。每个糖果的大小为 w1, w2, ... wm, 每个孩子对糖果大小的需求为
		n1, n2, ... nm。求怎样分糖果使得更多的孩子得到满足。
	贪婪思想： 将本问题抽象成给从 n 个孩子里选出 s 个孩子分糖果，让满足孩子的期望最大。贪婪的期望是使得 s 最大，
			贪婪的限制值是最多 m 个糖果。每次选出糖果需求最小的孩子，分配给他能满足的最小的糖果，以此类推
 */

func Candy(candy, children []int) int {
	selected := make([]bool, 0, len(children))
	allocted := make([]bool, 0, len(candy))
	ans := 0
	minNeed := math.MaxInt64
	for i := 0; i < len(children); i++ {
		if !selected[i] && children[i] < minNeed {
			minNeed = children[i]
		}
		temp := math.MaxInt64
		j := 0
		for ; j < len(candy); j++ {
			// 找出大于等于 minNeed 中最小的
			if !allocted[j] && candy[j] >= minNeed && candy[j] < temp {
				temp = candy[j]
			}
		}
		if temp != math.MaxInt64 {
			ans++
			selected[i] = true
			allocted[j] = true
		}
	}
	return ans
}