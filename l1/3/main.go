// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
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

	// WaitGroup это специальный счётчик, с помощью которого мы ожидаем
	// завершения работы всех go-рутин.
	wg := &sync.WaitGroup{}
	wg.Add(1)
	// Генерируем в канал входные данные.
	go Generator(wg, chIn)

	wg.Add(3)
	// Запускаем go-рутины.
	go Squaring(wg, chIn, chOut)
	go Squaring(wg, chIn, chOut)
	go Squaring(wg, chIn, chOut)

	go func() {
		wg.Wait()
		close(chOut)
	}()

	var sum int64
	for res := range chOut {
		sum += res
	}
	// Сумма квадратов.
	fmt.Println(sum)
}

func Generator(wg *sync.WaitGroup, chIn chan int64) {
	defer wg.Done()
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
