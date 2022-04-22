// Реализовать все возможные способы остановки выполнения go-рутины.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Способ 1 Отмена go-рутин по контексту.
	// Сюда же входят отмена по таймауту (context.WithTimeout),
	// по дедлайну (context.WithDeadline).
	// WithCancel возвращает ctx и функцию отмены.
	ctx, cancel := context.WithCancel(context.Background())
	
	// Запускаем go-рутину.
	go func(ctx context.Context)  {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Остановка go-рутины по контексу")
				return
			default:	
			}
			// 	Основная часть функции.
			time.Sleep(999*time.Millisecond)
			fmt.Println("Горутина (остановка по контексту) работает!")
		}
	}(ctx)

	time.Sleep(time.Second)
	// Вызываем функцию отмены, чтобы завершить go-рутину.
	cancel()

	// Прерывание go-рутины, через передачу сигнала в канал.
	// Создаём канал с пустой структурой, так как она не занимает места в памяти.
	ch := make(chan struct{})

	go func (ch chan struct{})  {
		for {
			select {
			case <-ch:
				fmt.Println("Остановка go-рутины по сигналу в переданному в канал.")
				return
			default:	
			}
			// 	Основная часть функции.
			time.Sleep(999*time.Millisecond)
			fmt.Println("Горутина (остановка по сигналу переданному в канал) работает!")
		}	
	}(ch)

	time.Sleep(time.Second)
	// Передаём в канал пустую структуру, чтобы завершить go-рутину.
	ch <- struct{}{}

	// Горутина завершиться, когда завершиться основная go-рутина main().
	go func()  {
		for {
			// 	Основная часть функции.
			time.Sleep(990*time.Millisecond)
			fmt.Println("Горутина работает!")
		}	
	}()

	time.Sleep(2*time.Second)
	fmt.Println("Основная go-рутина main() завершается! Все go-рутины будут остановлены!")
}