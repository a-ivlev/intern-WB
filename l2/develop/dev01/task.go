package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

// Функция, которая возвращает текущее сетевое время.
func ntpTime() (*time.Time, error) {
	netTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return nil, err
	}
	return &netTime, nil
}

func main() {

	netTimeReq, err := ntpTime()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(netTimeReq)
	fmt.Println(time.Now())
}
