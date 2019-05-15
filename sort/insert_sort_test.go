package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	nums := []int{10, 2, 31, 4, 51, 6, 17}
	nums1 := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println(InsertSort(nums))
	fmt.Println(InsertSort(nums1))
}
