package list

// 栈应用训练

// 表达式求值 3+3*9+4-9/3
func ExpressionEvaluation(ex string) (int, error)  {
	// 数字栈
	numStack := NewSliceStack(10)
	// 操作符栈
	operatorStack := NewSliceStack(10)
	for i := 0; i < len(ex); i++{
		c := ex[i]
		if 0 < c - '0' && c - '0' <= 9 {
			numStack.Push(int(c - '0'))
		}else {
			// 从操作符栈顶元素取元素
			if c == '+' || c == '-' {
				for ! operatorStack.IsEmpty() {
					num, _ := numStack.Pop()
					num1, _ := numStack.Pop()
					ost, _ := operatorStack.Pop()
					num2 := 0
					switch byte(ost) {
					case '+':
						num2 = num + num1
					case '-':
						num2 = num1 - num
					case '*':
						num2 = num1 * num
					case '/':
						num2 = num1 / num
					}
					numStack.Push(num2)
				}
				operatorStack.Push(int(c))
			}
			if c == '*' || c == '/' {
				for ! operatorStack.IsEmpty() {
					ost, _ := operatorStack.Top()
					if byte(ost) == '+' || byte(ost) == '-' {
						break
					}else {
						num, _ := numStack.Pop()
						num1, _ := numStack.Pop()
						num2 := 0
						operatorStack.Pop()
						switch byte(ost) {
						case '*':
							num2 = num1 * num
						case '/':
							num2 = num1 / num
						}
						numStack.Push(num2)
					}
				}
				operatorStack.Push(int(c))
			}
		}
	}

	for ! operatorStack.IsEmpty() {
		ost, _ :=  operatorStack.Pop()
		num, _ := numStack.Pop()
		num1, _ := numStack.Pop()
		num2 := 0
		switch byte(ost) {
		case '+':
			num2 = num + num1
		case '-':
			num2 = num1 - num
		case '*':
			num2 = num1 * num
		case '/':
			num2 = num1 / num
		}
		numStack.Push(num2)
	}
	return numStack.Top()
}

func ExpressionEvaluationII(ex string) int  {
	return -1
}