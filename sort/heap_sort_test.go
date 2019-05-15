package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	nums := []int{10, 2, 31, 4, 51, 6, 17}
	nums1 := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println(HeapSort(nums))
	fmt.Println(HeapSort(nums1))
}
