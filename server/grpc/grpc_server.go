package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/mushoffa/go-library/server"
	"google.golang.org/grpc"
)

// @Created 19/10/2021
// @Updated
type GrpcServer struct {
	Server *grpc.Server
	Port   int
}

// @Created 19/10/2021
// @Updated
func NewGrpcServer(port int, opt ...grpc.ServerOption) (server.ServerService, error) {

	grpcServer := grpc.NewServer(opt...)

	return &GrpcServer{grpcServer, port}, nil
}

// @Created 19/10/2021
// @Updated
func (s *GrpcServer) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errChannel := make(chan error, 1)
	signalChannel := make(chan os.Signal, 1)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		log.Fatalf("Failed to listen on %v", err)
		errChannel <- err
	}

	go func() {
		log.Printf("Grpc Server listening on port %v", listen.Addr())
		errChannel <- s.Server.Serve(listen)
	}()

	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChannel:
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

	case quit := <-signalChannel:
		log.Fatalf("signal.Notify: %v", quit)
		s.Server.GracefulStop()

	case done := <-ctx.Done():
		log.Fatalf("ctx.Done(): %v", done)
		s.Server.GracefulStop()
	}

	return nil
}
