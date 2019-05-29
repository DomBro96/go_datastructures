package list

import (
	"fmt"
	"testing"
)

func TestExpressionEvaluation(t *testing.T) {
	fmt.Println(ExpressionEvaluation("3+3*9+4-9/3"))
	fmt.Println(ExpressionEvaluation("3+4+5+6"))
}
