package graph

import (
	"testing"
)

func TestPartialSummation(t *testing.T) {
	nums := []int{1, 3, 5, 6, 10}
	t.Log(PartialSummation(nums, 15))
	t.Log(PartialSummation(nums, 11))
	t.Log(PartialSummation(nums, 1))
	t.Log(PartialSummation(nums, 20))
}

func TestFullPermutation(t *testing.T) {
	t.Logf("words: %v, full permutation: %v", []string{"hello", "world"}, FullPermutation([]string{"hello", "world"}))
	t.Logf("words: %v, full permutation: %v", []string{"hello", "world", "this", "is", "for", "you"}, FullPermutation([]string{"hello", "world", "this", "is", "for", "you"}))
}
