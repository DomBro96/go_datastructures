package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{23, 12, 78, 69, 100, 287, 89, 1, 5, 17, 32, 15, 90, 18, 2, 9, 11, 23, 9,
			 11, 11, 78, 8, 0, 987, 1024, 78, 6, 66, 15, 87, 90, 11, 200, 1456, 90, 80, 77, 254, 38}
	fmt.Println(len(nums))
	fmt.Println(QuickSort(nums))

}
