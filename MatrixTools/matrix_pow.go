package MatrixTools

import "math"

func Pow(matrix [][]float64, b float64) [][]float64 {
	result := matrix
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			result[i][j] = math.Pow(result[i][j], b)
		}
	}
	return result
}
