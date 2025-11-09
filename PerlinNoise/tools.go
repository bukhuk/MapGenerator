package perlin

import "math"

func fade(t float64) float64 {
	return 6*t*t*t*t*t - 15*t*t*t*t + 10*t*t*t
}

func inter(a, b, t float64) float64 {
	return a + t*(b-a)
}

var direction = [8][2]float64{
	{1 / math.Sqrt2, 1 / math.Sqrt2},
	{-1 / math.Sqrt2, 1 / math.Sqrt2},
	{1 / math.Sqrt2, -1 / math.Sqrt2},
	{-1 / math.Sqrt2, -1 / math.Sqrt2},
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func randomGradient(hash int, x, y float64) float64 {
	return x*direction[hash&7][0] + y*direction[hash&7][1]
}
