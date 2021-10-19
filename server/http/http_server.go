package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mushoffa/go-library/server"
)

// @Created 19/10/2021
// @Updated
type HttpServer struct {
	Port    int
	Handler http.Handler
}

// @Created 19/10/2021
// @Updated
func NewHttpServer(port int, handler http.Handler) server.ServerService {
	return &HttpServer{port, handler}
}

// @Created 19/10/2021
// @Updated
func (s *HttpServer) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errChannel := make(chan error, 1)
	signalChannel := make(chan os.Signal, 1)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: s.Handler,
	}

	go func() {
		log.Printf("Server is listening on port: ", s.Port)
		errChannel <- httpServer.ListenAndServe()
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
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	return nil
}
