package dp

import (
	"fmt"
	"testing"
)

func TestBagDp(t *testing.T) {
	fmt.Println(BagDp(5))
	fmt.Println(BagRec(0, 5))
	fmt.Println(BagDp2(5))
}
