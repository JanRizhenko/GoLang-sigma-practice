package main

func Min3(a, b, c float64) float64 {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func Average3(a, b, c float64) float64 {
	return (a + b + c) / 3.0
}

func SolveLinear(a, b float64) (float64, bool) {
	if a == 0 {
		return 0, false
	}
	return -b / a, true
}
