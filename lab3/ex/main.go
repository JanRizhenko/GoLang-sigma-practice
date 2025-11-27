package main

import (
	"fmt"
	"math"
)

const (
	a = uint64(134775813)
	c = uint64(1)
	m = uint64(1) << 32
	K = 50000
)

func GenerateIntSequence(seed uint64) []uint64 {
	x := seed
	seq := make([]uint64, K)

	for i := 0; i < K; i++ {
		x = (a*x + c) % m
		seq[i] = x % 100
	}
	return seq
}

func GenerateFloatSequence(seed uint64) []float64 {
	x := seed
	seq := make([]float64, K)

	for i := 0; i < K; i++ {
		x = (a*x + c) % m
		seq[i] = (float64(x) / float64(m)) * 100
	}
	return seq
}

func Analyze(seq []uint64) {
	freq := make([]int, 100)

	for _, v := range seq {
		freq[v]++
	}

	prob := make([]float64, 100)
	for i := 0; i < 100; i++ {
		prob[i] = float64(freq[i]) / float64(len(seq))
	}

	var mean float64
	for i := 0; i < 100; i++ {
		mean += float64(i) * prob[i]
	}

	var dispersion float64
	for i := 0; i < 100; i++ {
		dispersion += math.Pow(float64(i)-mean, 2) * prob[i]
	}

	std := math.Sqrt(dispersion)

	fmt.Println("=== Результати аналізу ===")
	fmt.Printf("Математичне сподівання: %.6f\n", mean)
	fmt.Printf("Дисперсія: %.6f\n", dispersion)
	fmt.Printf("СКВ: %.6f\n\n", std)

	fmt.Println("=== Ймовірності появи значень ===")
	for i := 0; i < 100; i++ {
		fmt.Printf("%2d : P = %.6f\n", i, prob[i])
	}
}

func main() {
	seq := GenerateIntSequence(1)
	Analyze(seq)

	floatSeq := GenerateFloatSequence(1)
	fmt.Println("Перші 5 дійсних значень [0;100):", floatSeq[:5])
}
