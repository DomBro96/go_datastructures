package greedy

import (
	"fmt"
	"testing"
)

func TestArmy(t *testing.T) {
	x := []int{1, 7, 15, 20, 30, 50}
	fmt.Println(Army(x, 10))
}