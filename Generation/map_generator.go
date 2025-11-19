package Generation

import (
	"MapGenerator/MatrixTools"
	perlin "MapGenerator/PerlinNoise"
	"MapGenerator/Teraforming"
	"MapGenerator/Visualization"
)

func Generate(n, m, seed int) {

	matrix := MatrixTools.Make(n, m, 0)

	scale := 256.
	k := 1.0

	for scale > 1. {
		p := perlin.NoiseMatrix(n, m, scale, seed)
		p.MultScalar(k)
		matrix.AddMatrix(p)
		k /= 1.5
		scale /= 2.
	}

	matrix.Norm()
	matrix.Pow(1.2)

	Map := Teraforming.Make(matrix)

	Map.ThermalErosion(16, 0.02, 0.5)
	Map.Blur(2)

	Map.HighRange(4, 64)

	Visualization.MakeOBJ(Map, "map.obj")
}
