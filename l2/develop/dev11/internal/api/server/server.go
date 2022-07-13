package server

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	srv http.Server
	wg  *sync.WaitGroup
}

func NewServer(addr string, h http.Handler) *Server {
	s := &Server{
		wg: &sync.WaitGroup{},
	}

	s.srv = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

func (s *Server) Start() {
	defer s.wg.Done()

	s.wg.Add(1)
	go func() {
		log.Printf("[ INFO ] Server started %s\n", s.srv.Addr)
		if err := s.srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	// Передаём контекст который прерывается по таймауту,
	// чтобы не зависнуть на какой-нибудь зависшей горутине.
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	cancel()
	s.wg.Wait()
}
