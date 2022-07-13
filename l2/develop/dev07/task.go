package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

// or принимает некоторое количество каналов, из которых можем только читать,
// возвращает один канал только для чтения.
func or(channels ...<-chan interface{}) <-chan interface{} {
	// Создаём главный канал с буфером равным числу переданных каналов.
	mergedCh := make(chan interface{}, len(channels))

	// Создаём WaitGroup для ожидания завершения работы go-рутин.
	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	// Итерируемся по переданным каналам и запускаем go-рутины.
	for _, channel := range channels {
		// Количество go-рутин, равно количеству переданных каналов.
		go func(chanel <-chan interface{}) {
			defer wg.Done()
			for ch := range chanel {
				mergedCh <- ch
			}
		}(channel)
	}

	// Ждём когда все каналы получат данные и закроются, после этого закрываем канал mergedCh.
	go func() {
		wg.Wait()
		close(mergedCh)
	}()

	return mergedCh
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(1*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
		sig(5*time.Second),

		//sig(2*time.Hour),
		//sig(5*time.Minute),
		//sig(1*time.Second),
		//sig(1*time.Hour),
		//sig(1*time.Minute),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}
