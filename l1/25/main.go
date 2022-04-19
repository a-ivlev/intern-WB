// Реализовать собственную функцию sleep.
//
package main

import (
	"fmt"
	"time"
)

func main() {
	Sleep(1 * time.Second)
	fmt.Println("Я проснулся!")
}

func Sleep(d time.Duration) {
	<-time.NewTimer(d).C
}