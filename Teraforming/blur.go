package Teraforming

func (cur *Map) Blur(passes int) {
	for step := 0; step < passes; step++ {
		for i := 0; i < cur.N; i++ {
			for j := 0; j < cur.M; j++ {
				sum := 0.0
				cnt := 0
				for addi := -1; addi < 2; addi++ {
					for addj := -1; addj < 2; addj++ {
						if 0 <= i+addi && i+addi < cur.N && 0 <= j+addj && j+addj < cur.M {
							sum += cur.A[i+addi][j+addj]
							cnt++
						}
					}
				}
				cur.A[i][j] = sum / float64(cnt)
			}
		}
	}
}
