// Реализовать паттерн «адаптер» на любом примере.
//
package main

import (
	"fmt"
	"log"
	"net/http"
)

// Самая часто встречаемая в Go реализация патерна адаптер
// это тип HandleFunc из пакета http.

// HandleFunc это тип функции имеющей методы и соответствующей интерфейсу http.Handler.
// HandleFunc является адаптером, который позволяет значению-функции соответствовать
// интерфейсу http.Handler. По сути, этот трюк позволяет типу database, соответствовать
// интерфейсу http.Handler несколькими различными способами: посредством его метода list,
// метода price и т.д.

type HandleFunc func(w http.ResponseWriter, r *http.Request)

func (f HandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

type database map[string]int64

func(db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %d\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	priceItem, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%d\n", priceItem)
}

func main() {
	db := database{
		"ботинки": 5000,
		"носки": 500,
	}

	mux := http.NewServeMux()
	mux.Handle("/list", HandleFunc(db.list))
	mux.Handle("/list", HandleFunc(db.list))
	// Поскольку регистрация обработчика таким образом является весьма распространённой,
	// ServeMux из пакета http имеет удобный метод HandleFunc:
	// HandleFunc registers the handler function for the given pattern.
	// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	// 	if handler == nil {
	// 		panic("http: nil handler")
	// 	}
	// 	mux.Handle(pattern, HandlerFunc(handler))
	// }

	log.Fatal(http.ListenAndServe(":8000", mux))
}