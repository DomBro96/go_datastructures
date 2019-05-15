package sort

import (
	"fmt"
	"testing"
)

func TestHellSort(t *testing.T) {
	nums := []int{10, 2, 31, 4, 51, 6, 17}
	nums1 := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println(HellSort(nums))
	fmt.Println(HellSort(nums1))

}
