package fibonacci

import "math"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func iround(x float64) int64 {
	return int64(math.Round(x))
}

func powi(x float64, p int) float64 {
	return math.Pow(x, float64(p))
}
