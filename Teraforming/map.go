package Teraforming

import "MapGenerator/MatrixTools"

type Map MatrixTools.Matrix

func Make(matrix MatrixTools.Matrix) Map {
	return Map{N: matrix.N, M: matrix.M, A: matrix.A}
}
