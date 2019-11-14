package greedy

/**
	栅栏问题： 切木板的开销 = 各个 叶子节点的深度*切割的开销
	找出最短，和次短的两块木块 l1, l2 计算开销，直到只剩下一块木块为止
 */

 /**
    length 为切割后的模板长度
  */
func FenceRepair(fence []int) int {
	n := len(fence)
	ans := 0
	// 只剩愉快木板时退出
	for n > 1 {
		// 寻找最短和次短木板
		mill0, mill1 := 0, 1
		if fence[mill0] > fence[mill1] {
			fence[mill0], fence[mill1] = fence[mill1], fence[mill0]
		}
		for i := 2; i < n; i++ {
			if fence[i] < fence[mill0] {
				mill1 = mill0
				mill0 = i
			}else if fence[i] < fence[mill1] {
				mill1 = i
			}
 		}


		t := fence[mill0] + fence[mill1]
		ans += t
		if mill0 == n - 1 {
			mill0, mill1 = mill1, mill0
		}
		fence[mill0] = t
		fence[mill1] = fence[n - 1]
		n--
	}
	return ans

}