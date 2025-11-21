package main

import (
	"fmt"
	"strconv"
)

func main() {
	variable8 := int8(127)
	variable16 := int16(16383)

	fmt.Println("Приведення типів\n")

	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n", uint8(variable16), uint8(variable16))

	// Завдання. //
	//1. Створіть 2 змінні різних типів. Виконайте арифметичні операції. Результат вивести в консоль
	variableAuto := 57
	variableInt16 := int16(27)
	variableString := "Check"

	// 1. Арифметика між сумісними типами — ОК
	sum1 := variableAuto + int(variableInt16)
	fmt.Println("\nРезультат додавання двох Int різних розрядів з перетворенням int16 на Int(auto):", sum1)

	// 2. Спроба додати int + string (імітація помилки)
	fmt.Println("\nСпроба додати int і string:")

	// Це неможливо → виводимо помилку самостійно
	fmt.Println("Помилка: неможливо додати 'int' і 'string'")

	// 3. Коректний варіант — конкатенація
	sum3 := strconv.Itoa(variableAuto) + variableString
	fmt.Println("\nРезультат конкатенації:", sum3)
}
