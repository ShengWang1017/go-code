package matrix

// SpiralOrder LC54
// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	m := len(matrix)    // 行数
	n := len(matrix[0]) // 列数

	left, right := 0, n-1
	top, bottom := 0, m-1

	result := make([]int, 0, m*n)

	// 保持“四边在同一循环体内”，但在第三、第四条边前做检查
	for left <= right && top <= bottom {
		// 1) 上边：固定 row = top, col 从 left -> right
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++ // 收缩上边界

		// 2) 右边：固定 col = right, row 从 top -> bottom
		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right-- // 收缩右边界

		// 3) 下边：在执行前检查 top <= bottom（避免只剩一行时重复）
		if top <= bottom {
			for col := right; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
			bottom-- // 收缩下边界
		}

		// 4) 左边：在执行前检查 left <= right（避免只剩一列时重复）
		if left <= right {
			for row := bottom; row >= top; row-- {
				result = append(result, matrix[row][left])
			}
			left++ // 收缩左边界
		}
	}

	return result
}
