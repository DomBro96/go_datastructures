package sort

import (
	"fmt"
	"testing"
)

func TestFindKthNum(t *testing.T) {
	nums := []int{4, 2, 5, 12, 3, 1, 15, 16, 9}
	// 3 2 4 12 5
	// 12 5
	fmt.Println(FindKthNum(nums, 1))
	fmt.Println(FindKthNum(nums, 3))
	fmt.Println(FindKthNum(nums, 5))
	fmt.Println(FindKthNum(nums, 9))
	fmt.Println(FindKthNum(nums, 7))


}
