package MatrixTools

import "math"

func ThermalErosion(matrix [][]float64, passes int, talus, k float64) [][]float64 {
	result := matrix
	for step := 0; step < passes; step++ {
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				for addi := -1; addi < 2; addi++ {
					for addj := -1; addj < 2; addj++ {
						if 0 <= i+addi && i+addi < len(matrix) && 0 <= j+addj && j+addj < len(matrix[i]) {
							if math.Abs(result[i][j]-result[i+addi][j+addj]) > talus {
								d := (result[i][j] - result[i+addi][j+addj] - talus) * k
								result[i][j] -= d
								result[i+addi][j+addj] += d
							}
						}
					}
				}
			}
		}
	}
	return result
}
