package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func RandomPermutation(seed int64) [512]int {
	var permutation [512]int
	r := rand.New(rand.NewSource(int64(seed)))
	base := r.Perm(256)
	for i := 0; i < 512; i++ {
		permutation[i] = base[i%256]
	}
	return permutation
}

func fade(t float64) float64 {
	return 6*t*t*t*t*t - 15*t*t*t*t + 10*t*t*t
}

func lerp(a, b, t float64) float64 {
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

func RandomGradient(hash int, x, y float64) float64 {
	return x*direction[hash&7][0] + y*direction[hash&7][1]
}

func Noise(x, y float64, perm [512]int) float64 {
	xi, yi := math.Floor(x), math.Floor(y)
	X, Y := int(xi)&255, int(yi)&255
	xf, yf := x-xi, y-yi

	h00 := perm[perm[X]+Y]
	h10 := perm[perm[X+1]+Y]
	h01 := perm[perm[X]+Y+1]
	h11 := perm[perm[X+1]+Y+1]

	corner00 := RandomGradient(h00, xf, yf)
	corner10 := RandomGradient(h10, xf-1, yf)
	corner01 := RandomGradient(h01, xf, yf-1)
	corner11 := RandomGradient(h11, xf-1, yf-1)

	u := fade(xf)
	v := fade(yf)

	return lerp(lerp(corner00, corner10, u), lerp(corner01, corner11, u), v)
}

func main() {
	seed, _ := strconv.Atoi(os.Args[1])
	if seed == 0 {
		seed = rand.Int()
	}
	w, h := 1024, 1024
	img := image.NewGray(image.Rect(0, 0, w, h))
	perm := RandomPermutation(int64(seed))
	scale := 100.0
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			h := Noise(float64(i)/scale, float64(j)/scale, perm)
			c := uint8((h + 1) * 256 * 0.5)
			img.SetGray(i, j, color.Gray{Y: c})
		}
	}
	f, _ := os.Create("perlin.png")
	defer f.Close()
	png.Encode(f, img)
}
