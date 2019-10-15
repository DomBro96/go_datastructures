package string_matching

import (
	"fmt"
	"testing"
)

func TestRKMatch(t *testing.T) {
	s := "abcdefghig"
	m := "hig"
	fmt.Println(RKMatch(s, m))
	s1 := "abcdefghig"
	m1 := "deghig"
	fmt.Println(RKMatch(s1, m1))
}
