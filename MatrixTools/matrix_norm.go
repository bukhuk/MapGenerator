package MatrixTools

func Norm(matrix [][]float64) [][]float64 {
	result := matrix
	mn := 1e9
	mx := -1e9
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			mn = min(mn, result[i][j])
			mx = max(mx, result[i][j])
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			result[i][j] = (result[i][j] - mn) / (mx - mn)
		}
	}
	return result
}
