// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала
// и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
//
// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	// В переменную worker праситься значение флага -wp, который определяет колличество
	// работающих воркеров.
	worker := flag.Int("wp", 1, "Define the number of workers, the default value is one.")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	
	defer cancel()

	chIn := make(chan int64)
	wg := sync.WaitGroup{}

	// Запуск определённого при запуске программы колличества воркеров.
	for i := 1; i <= *worker; i++ {
		wg.Add(1)
		go func(ctx context.Context, i int) {
			defer wg.Done()

			for x := range chIn {
				select {
				case <-ctx.Done():
					return
				default:
				}
				// Вывод в STDOUT информации какой воркер что прочитал.
				fmt.Printf("worker %d read = %d\n", i, x)
			}
		}(ctx, i)
	}

	// Постоянная запись данных в канал (главный поток).
	for x := int64(0); ; x++ {
		select {
			case <-ctx.Done():
				close(chIn)
				wg.Wait()
				return
			default:
			}
		chIn <- x
	}
}
