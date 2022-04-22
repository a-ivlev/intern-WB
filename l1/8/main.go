// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
package main

import "fmt"

func main() {
	n := int64(18)
	fmt.Printf("%08b\n", n)
	// Передаём в функцию ChangeOneBit переменную у которой необходимо изменить 1 бит.
	// Для изменения 1 бита на 1 вместе с переменной передаём true, для изменения 1 бита
	// на 0 передаём вместе с переменной false. В итоге получаем другое число.
	n = ChangeOneBit(n, false)

	fmt.Printf("%08b\n", n)
	fmt.Println(n)
}

// Для установки 1 биту значения 1, нужно указать znak = true,
// для установки 1 бита в значение 0, нужно указать znak = false.
func ChangeOneBit(n int64, znak bool) int64 {
	if znak {
		n = n | 1<<1
	}else{
		n = n&^(1 << 1)
	}
	return n
}