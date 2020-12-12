package array

// 一个二维数组，起每一行从左到右递增顺序排序，每列按从上到下递增顺序排序。完成这样一个函数，输入这样一个二维数组和
// 一个整数n 并判断二维数组中是否存在这个整数

// 解法一 暴力可解 时间复杂度 O(N*N)

// 解法二 通过判断矩阵右上角的数字 m , 其横坐标为x, 纵坐标为y。 如果 m > n 由于 m 已经是 y 列小的所以 x - 1
// 如果 m < n 由于 m 已经是 x 行最大的，所以 y + 1

func MatrixFind(matrix [][]int, num int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	row := 0
	col := cols - 1
	for row < rows && col >= 0 {
		m := matrix[row][col]
		if m > num {
			col--
		} else if m < num {
			row++
		} else {
			return true
		}
	}
	return false
}
