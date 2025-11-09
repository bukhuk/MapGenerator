package MatrixTools

func MultScalar(matrix [][]float64, l float64) [][]float64 {
	result := make([][]float64, len(matrix))
	for i := 0; i < len(matrix); i++ {
		result[i] = make([]float64, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			result[i][j] = matrix[i][j] * l
		}
	}
	return result
}
