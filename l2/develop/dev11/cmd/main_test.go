package main

import (
	"l2/develop/dev11/internal/api/handlers"
	"l2/develop/dev11/internal/app/event"
	"l2/develop/dev11/internal/store/inMemoryDB"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Тестируем хендлер /create_event насколько он корректно обрабатывает входящие запросы.
func TestRouter_CreateEvent(t *testing.T) {
	// Создаём реальный хендлер.
	store := inMemoryDB.NewInMemoryDB()
	eventStore := event.NewStore(store)
	rt := handlers.NewRouter(eventStore)
	//
	hts := httptest.NewServer(rt)

	var response, err = hts.Client().Post(hts.URL+"/create_event", "application/x-www-form-urlencoded", strings.NewReader("user_id=1&date=2019-09-09&description=test user_id 1"))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			t.Error(err)
		}
	}()
	if response.StatusCode != http.StatusCreated {
		t.Error("status wrong:", response.StatusCode)
	}
}
