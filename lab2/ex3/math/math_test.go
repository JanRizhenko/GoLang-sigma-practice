package math

import (
	"math"
	"testing"
)

func TestMin3(t *testing.T) {
	tests := []struct {
		name     string
		a, b, c  float64
		expected float64
	}{
		{"Positive numbers", 5, 2, 8, 2},
		{"Negative numbers", -1, -5, 3, -5},
		{"All equal", 7, 7, 7, 7},
		{"First is min", 1, 5, 10, 1},
		{"Last is min", 10, 5, 1, 1},
		{"Mixed signs", -2, 0, 5, -2},
		{"Zero values", 0, 0, 0, 0},
		{"Decimals", 3.5, 2.1, 8.9, 2.1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Min3(tt.a, tt.b, tt.c)
			if result != tt.expected {
				t.Errorf("Min3(%.2f, %.2f, %.2f) = %.2f; want %.2f",
					tt.a, tt.b, tt.c, result, tt.expected)
			}
		})
	}
}

func TestAverage3(t *testing.T) {
	tests := []struct {
		name     string
		a, b, c  float64
		expected float64
	}{
		{"Positive numbers", 10, 20, 30, 20},
		{"Same numbers", 5, 5, 5, 5},
		{"Negative numbers", -3, -6, -9, -6},
		{"Mixed signs", -5, 0, 5, 0},
		{"Zero values", 0, 0, 0, 0},
		{"Decimals", 1.5, 2.5, 3.0, 2.333333333333333},
		{"Large numbers", 100, 200, 300, 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Average3(tt.a, tt.b, tt.c)
			if math.Abs(result-tt.expected) > 1e-9 {
				t.Errorf("Average3(%.2f, %.2f, %.2f) = %.10f; want %.10f",
					tt.a, tt.b, tt.c, result, tt.expected)
			}
		})
	}
}

func TestSolveLinear(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
		hasRoot  bool
	}{
		{"Simple equation 2x + 6 = 0", 2, 6, -3, true},
		{"Simple equation 5x - 15 = 0", 5, -15, 3, true},
		{"Equation with a=0", 0, 5, 0, false},
		{"Negative coefficient", -3, 9, 3, true},
		{"Root is zero", 4, 0, 0, true},
		{"Decimal coefficients", 2.5, 5, -2, true},
		{"Large numbers", 100, 200, -2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := SolveLinear(tt.a, tt.b)
			if ok != tt.hasRoot {
				t.Errorf("SolveLinear(%.2f, %.2f) returned ok=%v; want ok=%v",
					tt.a, tt.b, ok, tt.hasRoot)
			}
			if ok && math.Abs(result-tt.expected) > 1e-9 {
				t.Errorf("SolveLinear(%.2f, %.2f) = %.10f; want %.10f",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func BenchmarkMin3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Min3(5.5, 2.3, 8.7)
	}
}

func BenchmarkAverage3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Average3(10, 20, 30)
	}
}

func BenchmarkSolveLinear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolveLinear(2, 6)
	}
}
