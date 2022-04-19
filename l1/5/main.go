// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	chData := make(chan string)

	// Создаём context с таймаутом.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Несмотря на то, что срок действия ctx истек, рекомендуется вызывать
	// функцию отмены в любом случае. Невыполнение этого требования может сохранить
	// контекст и его родителя дольше, чем необходимо. (Выдержка из https://pkg.go.dev/context#WithCancel)
	defer cancel()

	// Контекст с тайм-аутом, передаётся в функцию, чтобы сообщить функции, что она
	// должна прекратить свою работу по истечении времени ожидания.
	go func(ctx context.Context, chData chan string) {
		for i := 1; ; i++ {
			select{
			case <- ctx.Done():
				close(chData)
				return
			default:	
			}
			chData <- fmt.Sprintf("%d", i)
		}
	}(ctx, chData)

	for str := range chData {
		fmt.Println("Print", str)
	}
	fmt.Println("Program completed.")
}