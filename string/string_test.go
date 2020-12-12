package string

import (
	"fmt"
	"testing"
)

func TestReplaceBlank(t *testing.T) {
	s := "We Are Happy!"
	fmt.Println(ReplaceBlank(s))
	s1 := " "
	fmt.Println(ReplaceBlank(s1))
	s2 := "Whatdoyoudo"
	fmt.Println(ReplaceBlank(s2))

}
