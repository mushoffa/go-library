package http

import (
	"context"
	"fmt"
	"log"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"
)

type HttpServerService interface {
	Run()
}

type server struct {
	address 	string
	port 		int
	handler 	nethttp.Handler
}

func NewHttpServer(port int, handler nethttp.Handler) HttpServerService {
	return &server {
		address: fmt.Sprintf(":%d", port), 
		port: port,
		handler: handler,
	}
}

func (s *server) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errChannel := make(chan error, 1)
	signalChannel := make(chan os.Signal, 1)

	server := &nethttp.Server {
		Addr: s.address,
		Handler: s.handler,
	}

	go func() {
		log.Printf("Server is listening on port: ", s.port)
		errChannel <- server.ListenAndServe()
	}()

	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChannel:
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

	case quit := <-signalChannel:
		log.Fatalf("signal.Notify: %v", quit)
		cancel()

	case done := <-ctx.Done():
		log.Fatalf("ctx.Done(): %v", done)
	}

	// server.GracefulStop()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}