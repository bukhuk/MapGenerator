package MatrixTools

func Map(matrix [][]float64, low, high int) [][]int {
	matrix = Norm(matrix)
	result := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		result[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			result[i][j] = int(float64(low) + matrix[i][j]*float64(high-low))
		}
	}
	return result
}
