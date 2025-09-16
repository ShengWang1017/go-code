package matrix

// RotateImage
// 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
func RotateImage(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])
	layer := min(m, n)
	for i := 0; i < layer/2; i++ {

	}

}

// traverse
// 关键操作，遍历矩阵,一圈一圈
func traverse(matrix [][]int) []int {

	rows := len(matrix)
	colum := len(matrix[0])

	left := 0
	right := colum - 1
	top := 0
	bottom := rows - 1
	result := make([]int, rows*colum)
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		if left > right || bottom < top {
			break
		}
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])

		}
		bottom--

		if left > right || bottom < top {
			break
		}

		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])

		}
		left++

	}
	return result
}
