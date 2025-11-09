package main

import (
	"MapGenerator/MatrixTools"
	perlin "MapGenerator/PerlinNoise"
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	n, m := 512, 512
	matrix := make([][]float64, n)
	steps := 128
	for i := 0; i < n; i++ {
		matrix[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			matrix[i][j] = 0.
		}
	}

	scale := 256.
	k := 1.0

	for scale > 1. {
		matrix = MatrixTools.AddMatrix(matrix, MatrixTools.MultScalar(perlin.NoiseMatrix(n, m, scale, 0), k))
		k /= 1.5
		scale /= 2.
	}

	matrix = MatrixTools.Norm(matrix)

	matrix = MatrixTools.Pow(matrix, 1.2)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix[i][j] = math.Round(matrix[i][j]*float64(steps)) / float64(steps)
		}
	}

	matrix = MatrixTools.ThermalErosion(matrix, 16, 0.02, 0.5)

	matrix = MatrixTools.BoxBlur(matrix, 2)

	MatrixTools.ToIMG(matrix, "perlin.png")
	MatrixTools.ToOBJ(MatrixTools.Map(matrix, 4, 80), "map.obj")
	duration := time.Since(start)
	fmt.Printf("Time elapsed: %s\n", duration)
}
