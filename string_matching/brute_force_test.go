package string_matching

import (
	"fmt"
	"testing"
)

func TestBrouteForceMatch(t *testing.T) {
	s := "abcdefghig"
	m := "hig"
	fmt.Println(BruteForceMatch(s, m))
}
