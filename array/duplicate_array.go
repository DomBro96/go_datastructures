package array

import (
	"sort"
)

// 题目一 从长度为 n， 所有数字都在 0 ~ n-1 的数组内，找出任意一个重复数字。

// 方法一
// 长度为 n 所有数字都在 0 ~ n-1 的数组，在排序状态下数字 m 一定在下标为 m 的位置上；即 m = arr[m]
// 基于以上结论，可以遍历数组 m = arr[i], 当 m != i 时, 检查 arr[m] 是否等于 m, 若等于则说明 m
// 已经重复，返回 m; 若不等于则交换 m 和 arr[m], 此时 m 已经处于正确位置, 继续该步骤即可
// 时间复杂度 O(N), 空间复杂度 O(1)
func FindDupNumFromArray(nums []int) int {

	for i := range nums {
		m := nums[i]
		if m != i {
			if nums[m] == m {
				return m
			}
			nums[m], nums[i] = nums[i], nums[m]
		}
	}
	return -1
}

// 方法2
// 使用 hash 算法，根据 arr[i] % n 取 hash 值(此处由于数组中数字都小于n, 可以使用这么简单的 hash 函数)，
// 放入 hash table 中，若出现 hash 冲突说明该数字重复
// 时间复杂度 O(N), 空间复杂度 O(N)
func FindDupNumFromArrayWithHash(nums []int) int {

	hashTable := make([]bool, len(nums))
	for _, num := range nums {
		hansIndex := hash(num, len(nums))
		if hashTable[hansIndex] {
			return num
		}
		hashTable[hansIndex] = true
	}
	return -1
}

func hash(num int, length int) int {
	return num % length
}

// 方法三
// 对数组排序，遍历数组，发现 arr[i] == arr[i+1] 返回 arr[i]
// 排序的时间复杂度为 O(NlogN)
// 即总的时间复杂度也为 O(NlogN) 空间复杂度为 O(1) 综合来看还是方法一的性能最好

func FindDupNumFromArrayWithSort(nums []int) int {
	sort.Ints(nums)
	for i := range nums {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

// 题目二 不修改数组找出重复数字，长度为 n+1 的数组中所有的数字都在 1~n  的范围内，所以数组至少有一个数字是重复的。
// 请找出任意一个重复的数字，但不能修改输入的数组。

// 方法一
// 可以继续使用 hash 算法，借助 hash table 找出重复数字。
// 时间复杂度为 O(N) 空间复杂度为 O(N)

// 方法二
// 长度 n+1 数组所有数字都在 1~n 范围内，所以数组中必定存在一个重复的元素。可以在数组中寻找 1~2/n 的元素个数 m，
// 如果 m <= 2/n 说明在 1~2/n 可能没有重复数字。于是去剩下一半 2/n~n 的元素中查找个数 m1，该过程十分像二分查找
// 时间复杂度 O(NlogN) 空间复杂度 O(1)

func FindDupNumFfromArrayNoEdit(nums []int) int {
	startNum := 1
	endNum := len(nums) - 1
	for endNum >= startNum {
		midNum := ((endNum - startNum) >> 1) + startNum
		count := getNumCountByRange(startNum, midNum, nums)
		if startNum == endNum {
			if count > 1 {
				return startNum
			} else {
				break
			}
		}

		if count <= midNum-startNum+1 {
			startNum = midNum + 1
		} else {
			endNum = midNum
		}
	}

	return -1
}

func getNumCountByRange(startNum, endNum int, nums []int) int {
	count := 0
	for _, num := range nums {
		if num >= startNum && num <= endNum {
			count++
		}
	}
	return count
}
