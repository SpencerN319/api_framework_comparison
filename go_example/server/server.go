package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type builder struct {
	server      *http.Server
	gracePeriod time.Duration
	middlewares []Middleware
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func New(port string) *http.Server {
	return &http.Server{Addr: ":" + port}
}

func NewBuilder() *builder {
	return &builder{server: &http.Server{}}
}

func (b *builder) SetPort(port string) *builder {
	b.server.Addr = ":" + port
	return b
}

func (b *builder) SetReadTimeout(sec time.Duration) *builder {
	b.server.ReadTimeout = sec * time.Second
	return b
}

func (b *builder) SetWriteTimeout(sec time.Duration) *builder {
	b.server.WriteTimeout = sec * time.Second
	return b
}

func (b *builder) SetGracePeriod(sec time.Duration) *builder {
	b.gracePeriod = sec
	return b
}

func (b *builder) AddMiddleware(m Middleware) *builder {
	b.middlewares = append(b.middlewares, m)
	return b
}

func (b *builder) BuildAndServe() {
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// interrupt signal received, shut down.
		if err := b.server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := b.server.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("failed to initialize server: %v", err)
	}
	<-idleConnsClosed
}
