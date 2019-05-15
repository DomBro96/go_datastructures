package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{10, 2, 31, 4, 51, 6, 17}
	nums1 := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println(BubbleSort(nums))
	fmt.Println(BubbleSort(nums1))
	nums2 := []int{23, 12, 78, 69, 100, 287, 89, 1, 5, 17, 32, 15, 90, 18, 2, 9, 11, 23, 9,
		11, 11, 78, 8, 0, 987, 1024, 78, 6, 66, 15, 87, 90, 11, 200, 1456, 90, 80, 77, 254, 38}
	fmt.Println(BubbleSort(nums2))
}
