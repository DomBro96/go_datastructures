package graph

import (
	"fmt"
	"testing"
)

func TestPartialSummation(t *testing.T) {
	nums := []int{1, 3, 5, 6, 10}
	fmt.Println(PartialSummation(nums, 15))
	fmt.Println(PartialSummation(nums, 11))
	fmt.Println(PartialSummation(nums, 1))
	fmt.Println(PartialSummation(nums, 20))
}
