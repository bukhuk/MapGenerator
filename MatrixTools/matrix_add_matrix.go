package MatrixTools

func AddMatrix(matrixA, matrixB [][]float64) [][]float64 {
	result := make([][]float64, min(len(matrixA), len(matrixB)))
	for i := 0; i < min(len(matrixA), len(matrixB)); i++ {
		result[i] = make([]float64, min(len(matrixB[i]), len(matrixA[i])))
		for j := 0; j < min(len(matrixB[i]), len(matrixA[i])); j++ {
			result[i][j] = matrixA[i][j] + matrixB[i][j]
		}
	}
	return result
}
