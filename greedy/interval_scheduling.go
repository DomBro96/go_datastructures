package greedy

import "sort"

/**
	区间调度问题，给定每项工作开始时间 sn 和结束时间 en，ei > si 求不重合时间区间下(端点也不可以重合)，最多的可以
	完成多少工作。贪心思想，从剩余适合的选项中选择结束时间最早的
 */

type Interval struct {
	start int
	end int
}

func IntervalScheduling(n int, sn []int, en []int) int {
	intervals := make([]*Interval, 0, n)
	for i := 0; i < n; i++ {
		intervals = append(intervals, &Interval{sn[i], en[i]})
	}
	sort.Slice(intervals, func(i, j int) bool {
		return false
	})
	t, ans := 0, 0
	for i := 0; i < n; i++ {
		if t < intervals[i].start{
			ans++
			t = intervals[i].end
		}
	}
	return ans
}
