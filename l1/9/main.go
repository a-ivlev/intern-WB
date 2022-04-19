// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
package main

import "fmt"

func main() {
	// Созаём канал в который будут писатся числа из массива.
	chIn := make(chan int64)
	chOut := make(chan int64)

	array := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	// В этой go-рутине пишем числа из массива в канал.
	go func() {
		for _, x := range array {
			chIn <- x
		}
		close(chIn)
	}()

	// В этой go-рутине читаем данные из канала chIn, производим
	// метематическую операцию x*2 и результат операции пишем в канал chOut.
	go func() {
		for x := range chIn {
			chOut <- x * 2
		}
		close(chOut)
	}()

	// Читаем данные из канала chOut и выводим в stdout.
	for res := range chOut {
		fmt.Println(res)
	}
}