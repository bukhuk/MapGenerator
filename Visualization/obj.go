package Visualization

import (
	"MapGenerator/Teraforming"
	"fmt"
	"os"
)

func MakeOBJ(matrix Teraforming.Map, path string) {
	f, _ := os.Create(path)
	defer f.Close()

	mapPoints := 0

	for i := 0; i < matrix.N; i++ {
		for j := 0; j < matrix.M; j++ {
			fmt.Fprintln(f, "v", i, j, matrix.A[i][j])
			mapPoints++
		}
	}

	for i := 0; i < matrix.N; i++ {
		fmt.Fprintln(f, "v", i, 0, 0)
	}

	for i := 0; i < matrix.N; i++ {
		fmt.Fprintln(f, "v", i, matrix.M-1, 0)
	}

	for i := 0; i < matrix.M; i++ {
		fmt.Fprintln(f, "v", 0, i, 0)
	}

	for i := 0; i < matrix.M; i++ {
		fmt.Fprintln(f, "v", matrix.N-1, i, 0)
	}

	fmt.Fprintln(f)
	for i := 0; i < matrix.N-1; i++ {
		for j := 0; j < matrix.M-1; j++ {
			v00 := i*matrix.M + j + 1
			v01 := i*matrix.M + j + 2
			v10 := (i+1)*matrix.M + j + 1
			v11 := (i+1)*matrix.M + j + 2
			fmt.Fprintln(f, "f", v00, v01, v11)
			fmt.Fprintln(f, "f", v00, v10, v11)
		}
	}

	for i := 0; i < matrix.N-1; i++ {
		v00 := mapPoints + i + 1
		v01 := mapPoints + i + 2
		v10 := i*matrix.M + 1
		v11 := (i+1)*matrix.M + 1
		fmt.Fprintln(f, "f", v00, v01, v11)
		fmt.Fprintln(f, "f", v00, v10, v11)
	}

	for i := 0; i < matrix.N-1; i++ {
		v00 := mapPoints + matrix.N + i + 1
		v01 := mapPoints + matrix.N + i + 2
		v10 := (i + 1) * matrix.M
		v11 := (i + 2) * matrix.M
		fmt.Fprintln(f, "f", v00, v01, v11)
		fmt.Fprintln(f, "f", v00, v10, v11)
	}

	for i := 0; i < matrix.M-1; i++ {
		v00 := mapPoints + 2*matrix.N + i + 1
		v01 := mapPoints + 2*matrix.N + i + 2
		v10 := i + 1
		v11 := i + 2
		fmt.Fprintln(f, "f", v00, v01, v11)
		fmt.Fprintln(f, "f", v00, v10, v11)
	}

	for i := 0; i < matrix.M-1; i++ {
		v00 := mapPoints + 2*matrix.N + matrix.M + i + 1
		v01 := mapPoints + 2*matrix.N + matrix.M + i + 2
		v10 := (matrix.N-1)*matrix.M + i + 1
		v11 := (matrix.N-1)*matrix.M + i + 2
		fmt.Fprintln(f, "f", v00, v01, v11)
		fmt.Fprintln(f, "f", v00, v10, v11)
	}
	fmt.Fprintln(f, "f", mapPoints+1, mapPoints+matrix.N, mapPoints+2*matrix.N+2*matrix.M)
	fmt.Fprintln(f, "f", mapPoints+1, mapPoints+2*matrix.N+matrix.M, mapPoints+2*matrix.N+2*matrix.M)
}
