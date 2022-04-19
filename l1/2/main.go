// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаём 2 канала.
	// Канал для входных данных.
	chIn := make(chan int64)
	// Канал для результирующих данных.
	chOut := make(chan int64)

	// Генерируем в канал входные данные.
	go Generator(chIn)

	// WaitGroup это специальный счётчик, с помощью которого мы ожидаем
	// завершения работы всех go-рутин.
	wg := &sync.WaitGroup{}
	wg.Add(3)
	// Запускаем go-рутины.
	go Squaring(wg, chIn, chOut)
	go Squaring(wg, chIn, chOut)
	go Squaring(wg, chIn, chOut)

	go func() {
		wg.Wait()
		close(chOut)
	}()

	for res := range chOut {
		fmt.Println(res)
	}
	// В STDOUT получим:
	// 4
	// 64
	// 100
	// 16
	// 36
	// Последовательность может быть любой,
	// т.к. вычисления производятся конкурентно.
}

func Generator(chIn chan int64) {
	// Исходные данные.
	array := []int64{2, 4, 6, 8, 10}
	// Отправка исходных данных в канал.
	for _, elem := range array {
		chIn <- elem
	}
	// Мы отправили в канал все данные и после этого необходимо закрыть канал.
	// Чтобы получающие go-подпрограммы могли прекратить ожидание. Иначе будет deadlock!
	close(chIn)
}

func Squaring(wg *sync.WaitGroup, chIn, chOut chan int64) {
	defer wg.Done()
	for x := range chIn {
		chOut <- x * x
	}
}
