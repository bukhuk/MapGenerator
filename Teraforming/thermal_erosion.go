package Teraforming

import "math"

func (cur *Map) ThermalErosion(passes int, talus, k float64) {
	for step := 0; step < passes; step++ {
		for i := 0; i < cur.N; i++ {
			for j := 0; j < cur.M; j++ {
				for addi := -1; addi < 2; addi++ {
					for addj := -1; addj < 2; addj++ {
						if 0 <= i+addi && i+addi < cur.N && 0 <= j+addj && j+addj < cur.M {
							if math.Abs(cur.A[i][j]-cur.A[i+addi][j+addj]) > talus {
								d := (cur.A[i][j] - cur.A[i+addi][j+addj] - talus) * k
								cur.A[i][j] -= d
								cur.A[i+addi][j+addj] += d
							}
						}
					}
				}
			}
		}
	}
}
