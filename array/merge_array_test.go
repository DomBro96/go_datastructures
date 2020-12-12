package array

import (
	"fmt"
	"testing"
)

func TestMergeArray(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{6, 7, 8, 9, 10}
	fmt.Println(MergeArray(arr1, arr2))
	arr3 := []int{1, 3, 5, 7, 9}
	arr4 := []int{2, 4, 6, 8, 10}
	fmt.Println(MergeArray(arr3, arr4))
}
