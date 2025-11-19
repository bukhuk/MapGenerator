package Teraforming

func (cur *Map) HighRange(low, high float64) {
	mn := 1e9
	mx := -1e9
	for i := 0; i < cur.N; i++ {
		for j := 0; j < cur.M; j++ {
			mn = min(mn, cur.A[i][j])
			mx = max(mx, cur.A[i][j])
		}
	}
	for i := 0; i < cur.N; i++ {
		for j := 0; j < cur.M; j++ {
			cur.A[i][j] = (cur.A[i][j] - mn) / (mx - mn)
			cur.A[i][j] = low + cur.A[i][j]*(high-low)
		}
	}
}
