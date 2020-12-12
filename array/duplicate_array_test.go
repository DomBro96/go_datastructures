package array

import (
	"fmt"
	"testing"
)

func TestFindDupNumFromArray(t *testing.T) {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(FindDupNumFromArray(nums))
	nums1 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(FindDupNumFromArray(nums1))
	nums2 := []int{1, 2, 3, 0, 0, 0}
	fmt.Println(FindDupNumFromArray(nums2))
}

func TestFindDupNumFromArrayWithHash(t *testing.T) {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(FindDupNumFromArrayWithHash(nums))
	nums1 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(FindDupNumFromArrayWithHash(nums1))
	nums2 := []int{1, 2, 3, 0, 0, 0}
	fmt.Println(FindDupNumFromArrayWithHash(nums2))
}

func TestFindDupNumFromArrayNoEdit(t *testing.T) {
	nums := []int{2, 3, 5, 4, 3, 2, 6, 7}
	fmt.Println(FindDupNumFromArrayNoEdit(nums))
	nums1 := []int{2, 2, 5, 4, 2, 2, 6, 7}
	fmt.Println(FindDupNumFromArrayNoEdit(nums1))
	nums2 := []int{1, 2, 3, 4, 5, 6, 6, 7}
	fmt.Println(FindDupNumFromArrayNoEdit(nums2))
}
