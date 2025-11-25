package main

import "fmt"

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

func main() {
	fmt.Println("=== Тестування Min3 ===")
	fmt.Printf("Min3(5, 2, 8) = %.2f\n", Min3(5, 2, 8))
	fmt.Printf("Min3(-1, -5, 3) = %.2f\n", Min3(-1, -5, 3))

	fmt.Println("\n=== Тестування Average3 ===")
	fmt.Printf("Average3(10, 20, 30) = %.2f\n", Average3(10, 20, 30))
	fmt.Printf("Average3(5, 15, 25) = %.2f\n", Average3(5, 15, 25))

	fmt.Println("\n=== Тестування SolveLinear ===")
	if x, ok := SolveLinear(2, 6); ok {
		fmt.Printf("Рівняння 2x + 6 = 0, x = %.2f\n", x)
	} else {
		fmt.Println("Рівняння не має розв'язку")
	}

	if x, ok := SolveLinear(5, -15); ok {
		fmt.Printf("Рівняння 5x - 15 = 0, x = %.2f\n", x)
	} else {
		fmt.Println("Рівняння не має розв'язку")
	}

	if x, ok := SolveLinear(0, 5); ok {
		fmt.Printf("Рівняння 0x + 5 = 0, x = %.2f\n", x)
	} else {
		fmt.Println("Рівняння 0x + 5 = 0 не має розв'язку")
	}
}
