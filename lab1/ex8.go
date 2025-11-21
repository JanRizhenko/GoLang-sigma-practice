package main

//Импорт нескольких пакетов
import (
	"fmt"
	"math"
)

func main() {
	var defaultFloat float32
	var defaultDouble float64 = 5.5

	fmt.Println("defaultfloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")

	// Завдання.
	// 1. Створіть змінні різних типів, використовуючи короткий запис та ініціалізацію за замовчуванням. Результат вивести в консоль
	intVar := 42
	floatVar := 3.14159
	boolVar := true
	stringVar := "Hello, Go!"
	complexVar := complex(2, 3)

	fmt.Println("Змінні різних типів:")
	fmt.Printf("intVar       = %v (%T)\n", intVar, intVar)
	fmt.Printf("floatVar     = %v (%T)\n", floatVar, floatVar)
	fmt.Printf("boolVar      = %v (%T)\n", boolVar, boolVar)
	fmt.Printf("stringVar    = %v (%T)\n", stringVar, stringVar)
	fmt.Printf("complexVar   = %v (%T)\n", complexVar, complexVar)
}
