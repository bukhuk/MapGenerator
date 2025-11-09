package perlin

import "math/rand"

func randomPermutation(seed int64) [512]int {
	var permutation [512]int
	r := rand.New(rand.NewSource(seed))
	base := r.Perm(256)
	for i := 0; i < 512; i++ {
		permutation[i] = base[i%256]
	}
	return permutation
}
