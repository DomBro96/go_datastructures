package backtracking

import "fmt"

// 回溯算法经典例题 八皇后

var (
	// 标记皇后在 8x8 棋盘中位置, 索引表示所在行, 值表示所在列，初始状态8个皇后在每行的第一列上
	queens = make([]int, 8)
	result = 0
)

func EightQueens(row int) {
	if row == 8 {
		result++
		fmt.Println(result)
		PrintEightQueens()
		return
	}
	// 按列判断
	for column := 0; column < 8; column++ {
		if isOk(row, column) {
			queens[row] = column
			// 继续下一行
			EightQueens(row + 1)
		}
	}
}

// 皇后放到这一列是否合法
func isOk(row, column int) bool {
	// 左上角位置
	upperLeft := column - 1
	// 右上角位置
	upperRight := column + 1
	for r := row - 1; r >= 0; r-- {
		// 上面位置
		if queens[r] == column {
			return false
		}
		if queens[r] == upperLeft {
			return false
		}
		if queens[r] == upperRight {
			return false
		}
		upperLeft--
		upperRight++
	}
	return true
}

func PrintEightQueens() {
	for row := 0; row < 8; row++ {
		queenColumn := queens[row]
		for column := 0; column < 8; column++ {
			if column != queenColumn {
				fmt.Print(". ")
			} else {
				fmt.Print("Q ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
