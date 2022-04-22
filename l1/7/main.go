// Реализовать конкурентную запись данных в map.
package main

import (
	"fmt"
	"strconv"
	"sync"
)

// createMap осуществляет конкурентную запись в map используя мьютекс.
func createMap(mu *sync.Mutex, mapDB map[int]string, key int, value string) {
	mu.Lock()
	mapDB[key] = value
	mu.Unlock()
}


func main() {
	mapDB := make(map[int]string, 100)
	mu := &sync.Mutex{}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(wg *sync.WaitGroup)  {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			createMap(mu, mapDB, i, strconv.Itoa(i))
		}		
	}(wg)

	go func(wg *sync.WaitGroup)  {
		defer wg.Done()
		for i := 5; i < 10; i++ {
			createMap(mu, mapDB, i, strconv.Itoa(i))
		}
	}(wg)

	wg.Wait()

	fmt.Println(mapDB)
}

