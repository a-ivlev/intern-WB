package main

import "fmt"

/* Дан фрагмент кода.

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	justString = v[:100]
}

func main() {
	someFunc()
}

*/

// Использование глобальной переменной приведёт к тому,
// что огромная строка хранящаяся в переменной v будет храниться в памяти
// и не будет очищена GC. Так как мы часть этой строки храним в глобальной переменной.
// Если нам необходим результат работы функции мы должны возвращать его из функции,
// а не присваивать глобальной переменной.

// Исправленный вариант
func someFunc() string {
	v := createHugeString(1 << 10)

	return v[:100]
}

// Для разбора данного примера сама реализация неважна.
func createHugeString(size int) string{
	bs := make([]byte, size)
	bl := 0
	for i := 0; i < size; i++ {
		bl += copy(bs[bl:], []byte("Z"))
	}
	
	return string(bs[:])
}

func main() {
	justString := someFunc()

	fmt.Println(justString)
}
