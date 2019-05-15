package heap

// 和堆有关的练习
// 汽车从 起点a 开往 终点b， 路过 N 个加油站， 数组 A[N] 中 A[i] 表示 加油站i 距离起点a 的距离，
// 数组 B[N] 中 B[i] 表示 加油站i 最多可以加的油。设汽车每行驶 1 单位的距离，就要消耗 1 单位的汽油，
// 初始汽油为 P，a 到 b 的距离为 L， 且汽车油箱容量无限，求汽车行驶过程中最少的加油次数。

func PetrolStation(N, L, P int, A, B []int) int {
	// 0. 到达一个加油站，视为在下一个加油站拥有了在该加油站的权利。
	// 1. 每经过一个加油站，将其加入优先队列，直到汽车没油时从优先队列中pop出可以加的最大的加油量(第0条的原则)，
	//    选用优先队列是因为要求最小加油次数。
	// 2. 直到从优先队列中 pop 出的加油数量够行驶到下一个加油站。
	A = append(A, 0)
	B = append(B, L)
	pq := NewPriorityQueue(N+1)
	res, tank, pos := 0, P, 0
	for i := 0; i < N; i++ {
		// 要前进的距离
		dis := A[i] - pos
		for tank - dis < 0 {
			if pq.size == 0 {
				return -1
			}
			tank = pq.Pop()
			res++
		}
		tank -= dis
		pos = A[i]
		pq.Push(A[i])
	}

	return res


}
