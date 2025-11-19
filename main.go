package main

import (
	"MapGenerator/MatrixTools"
	perlin "MapGenerator/PerlinNoise"
	"MapGenerator/Visualization"
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()

	n, m := 256, 512
	matrix := MatrixTools.Make(n, m, 0)
	steps := 128

	scale := 256.
	k := 1.0

	for scale > 1. {
		p := perlin.NoiseMatrix(n, m, scale, 0)
		p.MultScalar(k)
		matrix.AddMatrix(p)
		k /= 1.5
		scale /= 2.
	}

	matrix.Norm()
	matrix.Pow(1.2)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix.A[i][j] = math.Round(matrix.A[i][j]*float64(steps)) / float64(steps)
		}
	}

	matrix.ThermalErosion(16, 0.02, 0.5)
	matrix.BoxBlur(2)

	Visualization.MakeIMG(matrix, "perlin.png")

	matrix.Map(4, 64)

	Visualization.MakeOBJ(matrix, "map.obj")

	duration := time.Since(start)
	fmt.Printf("Time elapsed: %s\n", duration)
}
