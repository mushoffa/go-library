package grpc

import (
	"context"

	"google.golang.org/grpc"
)

// @Created 19/10/2021
// @Updated
type GrpcClient struct {
	Conn *grpc.ClientConn
}

// @Created 19/10/2021
// @Updated
func NewGrpcClient(url string, opts ...grpc.DialOption) (*GrpcClient, error) {
	conn, err := grpc.Dial(url, opts...)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	return &GrpcClient{conn}, nil
}

// @Created 19/10/2021
// @Updated
func NewGrpcClientWithContext(url string, opts ...grpc.DialOption) (*GrpcClient, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.DialContext(ctx, url, opts...)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	return &GrpcClient{conn}, nil
}
