package MatrixTools

func BoxBlur(matrix [][]float64, passes int) [][]float64 {
	result := matrix
	for step := 0; step < passes; step++ {
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				sum := 0.0
				cnt := 0
				for addi := -1; addi < 2; addi++ {
					for addj := -1; addj < 2; addj++ {
						if 0 <= i+addi && i+addi < len(matrix) && 0 <= j+addj && j+addj < len(matrix[i]) {
							sum += matrix[i+addi][j+addj]
							cnt++
						}
					}
				}
				result[i][j] = sum / float64(cnt)
			}
		}
	}
	return result
}
