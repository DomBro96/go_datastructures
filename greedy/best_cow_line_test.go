package greedy

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := "abcdefg"
	fmt.Println(ReverseString(s))
	s1 := "abcdef"
	fmt.Println(ReverseString(s1))
}

func TestBestCowLine(t *testing.T) {
	fmt.Println(BestCowLine("ACDBCB"))
}