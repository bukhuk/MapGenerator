package perlin

import "math"

func noise(x, y float64, perm [512]int) float64 {
	xi, yi := math.Floor(x), math.Floor(y)
	X, Y := int(xi)&255, int(yi)&255
	xf, yf := x-xi, y-yi

	h00 := perm[perm[X]+Y]
	h10 := perm[perm[X+1]+Y]
	h01 := perm[perm[X]+Y+1]
	h11 := perm[perm[X+1]+Y+1]

	corner00 := randomGradient(h00, xf, yf)
	corner10 := randomGradient(h10, xf-1, yf)
	corner01 := randomGradient(h01, xf, yf-1)
	corner11 := randomGradient(h11, xf-1, yf-1)

	u := fade(xf)
	v := fade(yf)

	return inter(inter(corner00, corner10, u), inter(corner01, corner11, u), v)
}
