// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"l1/18/counter"
)

func main() {
	count := counter.NewCount()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func (wg *sync.WaitGroup) {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			count.Inc()
		}
	}(wg)

	go func (wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			count.Inc()
		}

	}(wg)

	wg.Wait()

	fmt.Println("Счетчик =", count.Get())

}