package main

import "fmt"

func main() {
	// Створення булевих змінних
	var first, second bool // за замовчуванням false
	var third bool = true  // явно true
	fourth := !third       // NOT third → false
	var fifth = true       // короткий запис

	fmt.Println("Булеві змінні:")
	fmt.Println("first  = ", first, " // false, за замовчуванням")
	fmt.Println("second = ", second, " // false, за замовчуванням")
	fmt.Println("third  = ", third, " // true, явно задано")
	fmt.Println("fourth = ", fourth, " // !true → false")
	fmt.Println("fifth  = ", fifth, " // true, короткий запис\n")

	// Логічне заперечення (!)
	fmt.Println("Заперечення (!):")
	fmt.Println("!true  = ", !true, " // false")
	fmt.Println("!false = ", !false, " // true\n")

	// Логічне AND (&&)
	fmt.Println("Логічне AND (&&):")
	fmt.Println("true && true   = ", true && true, " // обидва true → true")
	fmt.Println("true && false  = ", true && false, " // один false → false")
	fmt.Println("false && false = ", false && false, " // обидва false → false\n")

	// Логічне OR (||)
	fmt.Println("Логічне OR (||):")
	fmt.Println("true || true   = ", true || true, " // хоча б один true → true")
	fmt.Println("true || false  = ", true || false, " // хоча б один true → true")
	fmt.Println("false || false = ", false || false, " // обидва false → false\n")

	// Порівняння чисел
	fmt.Println("Порівняння чисел:")
	fmt.Println("2 < 3  = ", 2 < 3, " // 2 менше 3 → true")
	fmt.Println("2 > 3  = ", 2 > 3, " // 2 не більше 3 → false")
	fmt.Println("3 < 3  = ", 3 < 3, " // 3 не менше 3 → false")
	fmt.Println("3 <= 3 = ", 3 <= 3, " // 3 менше або дорівнює 3 → true")
	fmt.Println("3 > 3  = ", 3 > 3, " // 3 не більше 3 → false")
	fmt.Println("3 >= 3 = ", 3 >= 3, " // 3 більше або дорівнює 3 → true")
	fmt.Println("2 == 3 = ", 2 == 3, " // 2 не дорівнює 3 → false")
	fmt.Println("3 == 3 = ", 3 == 3, " // 3 дорівнює 3 → true")
	fmt.Println("2 != 3 = ", 2 != 3, " // 2 не дорівнює 3 → true")
	fmt.Println("3 != 3 = ", 3 != 3, " // 3 дорівнює 3 → false")

	// Завдання.
	// 1. Пояснити результати операцій (виконано)
}
