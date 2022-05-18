package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// HttpDoer необходим для теста.
type HttpDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client создаём http клиент. В структуру можно передать таймаут, куки и прочую информацию о запросе.
type Client struct {
	http HttpDoer
}

// NewClient конструктор, создаёт клиента по умолчанию.
func NewClient() *Client {
	return &Client{
		http: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// SetHttp в этот метод можно передать любую реализацию интерфейса HttpDoer.
func (c *Client) SetHttp(doer HttpDoer) *Client {
	c.http = doer
	return c
}

// getURL отправляет запрос к удалённому серверу, возвращает слайс байт и ошибку.
func (c Client) getURL(url string) ([]byte, error) {
	//// Создаём http клиент. В структуру можно передать таймаут, куки и прочую информацию о запросе.
	//client := http.Client{}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		fmt.Println("URL transformed to HTTPS due to an HSTS policy")
		url = "https://" + url
	}
	fmt.Println(url)

	fmt.Print("HTTP-запрос отправлен. ")
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// нужно закрывать тело, когда прочитаем что нужно.
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	fmt.Print("Ожидание ответа… ")

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.Header.Get("Date"))
	fmt.Println(resp.Header.Get("Content-Length"))

	return content, nil
}

func main() {
	output := flag.String("o", "index.html", "number of lines to read from the file")
	flag.Parse()

	client := NewClient()

	content, err := client.getURL(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Сохранение в файл: %s\n", *output)

	err = ioutil.WriteFile(*output, content, 0644)
	if err != nil {
		log.Fatalln("WriteFile: ", err)
	}
}
