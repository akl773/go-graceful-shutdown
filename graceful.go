package graceful

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

// Server is a wrapper for the http Server
type Server struct {
	httpServer *http.Server
}

func NewGracefulServer(addr string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (s *Server) Start() error {
	done := make(chan bool, 1)
	go s.handleSignals(done)
	err := s.httpServer.ListenAndServe()
	<-done
	return err
}

func (s *Server) handleSignals(done chan<- bool) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, os.Kill)

	<-sigint

	// We received an interrupt/kill signal, shut down.
	if err := s.Shutdown(context.Background()); err != nil {
		// Error from closing listeners, or context timeout:
		fmt.Printf("HTTP server Shutdown: %v", err)
	}
	done <- true
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
