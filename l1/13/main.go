// Поменять местами два числа без создания временной переменной.
package main

import "fmt"

func main() {
	a := 5
	b := 10

	// 1 Способ.
	a, b = b, a
	fmt.Printf("a = %d b = %d\n", a, b)


	// 2 Способ.
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a = %d b = %d\n", a, b)

}