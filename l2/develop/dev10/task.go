package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	var timeOut time.Duration

	flag.DurationVar(&timeOut, "timeout", 10, "таймаут")
	flag.Parse()

	if len(flag.Args()) < 2 {
		log.Fatal("Не указан хост и порт для подключения.")
	}

	host := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	// Конфигурируем подключение.
	dialer := net.Dialer{
		KeepAlive: time.Second * timeOut,
		Timeout:   time.Second * timeOut,
	}

	conn, err := dialer.DialContext(ctx, "tcp", host)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	go func() {
		defer conn.Close()
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Println(err)
		}
		// Функция отмены срабатывает при закрытии соединения со стороны сервера.
		cancel()
		return
	}()

	go func() {
		if _, err := io.Copy(conn, os.Stdin); err != nil {
			log.Println(err)
		}
		// Функция отмены срабатывает при нажатии Ctrl+D
		cancel()
	}()

	<-ctx.Done()

	fmt.Printf("%s: exit\n", conn.LocalAddr())
}
