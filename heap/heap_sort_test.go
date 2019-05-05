package heap

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1, 8, 9, 6, 10}
	HeapSort(nums)
	fmt.Println(nums)
}