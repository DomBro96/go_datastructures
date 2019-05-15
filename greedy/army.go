package greedy


// 直线上有 N 个点，点i的位置是 Xi，在 N 个点中寻找若干个点，给他们加上标记，其距离为R以内的区域里必须带有
// 标记的点。求至少有多少个点被加上标记。
// 贪婪算法： 从最左边的点 left 开始，寻找在 R 范围内距离 left 最大的点。对于第一个超过 R 的范围的点采用同样的算法
//
func Army(x []int, r int) int {
	i, ans := 0, 0
	for i < len(x) {
		// s 没有被覆盖的最左的点
		s := x[i]
		i++
		// 一直向右， 直到 s 右边第一个范围在 R 以外的点。
		for i < len(x) && x[i] <= s + r {
			i++
		}
		// p 标记点
		p := x[i-1]
		// 一直向右，直到 p 右边第一个范围在 R 以外的点。
		for i < len(x) && x[i] <= p + r {
			i++
		}
		ans++
	}
	return ans

}
