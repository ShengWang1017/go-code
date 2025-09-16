package matrix

// SetZeroes LC73. 矩阵置零
// 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
func SetZeroes(matrix [][]int) [][]int {
	row := len(matrix)
	colum := len(matrix[0])
	rowZero := make([]int, row)
	columZero := make([]int, colum)

	for i := 0; i < row; i++ {
		for j := 0; j < colum; j++ {
			if matrix[i][j] == 0 {
				rowZero[i] = -1
				columZero[j] = -1
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < colum; j++ {
			if rowZero[i] == -1 || columZero[j] == -1 {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}
