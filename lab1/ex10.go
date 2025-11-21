package main

import "fmt"

func main() {
	// Створюємо змінну типу int8 з символом 'R'
	var chartype int8 = 'R'
	fmt.Printf("Code '%c' - %d\n", chartype, chartype)
	// Виводимо: Code 'R' - 82
	// 82 — це код символу 'R' в ASCII

	// 1. Виведення української літери 'Ї'
	var ukrainianRune rune = 'Ї'
	fmt.Printf("Code '%c' - %d\n", ukrainianRune, ukrainianRune)
	// Виведе: Code 'Ї' - 1025
	// 1025 — це Unicode код для 'Ї'

	// 2. Пояснення типу "rune"
	/*
	   rune — це int32, який використовується для зберігання одного
	   Unicode-символу. Це дозволяє працювати з літерами будь-яких мов,
	   включно з українськими, китайськими, арабськими тощо.
	   Змінні типу rune зручно застосовувати для роботи з текстом у форматі Unicode.
	*/
}
