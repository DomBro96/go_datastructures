package greedy

import (
	"fmt"
	"testing"
)

func TestCoinSolve(t *testing.T) {
	num := []int{3, 2, 1, 3, 0, 2}
	fmt.Println(CoinSolve(num, 620))
}
