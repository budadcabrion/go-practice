package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/budadcabrion/go-practice/service"

	"google.golang.org/grpc"
)

const (
	port = ":12345"
)

type serviceServer struct {
	service.UnimplementedServiceServer
}

func (s *serviceServer) Time(ctx context.Context, in *service.TimeRequest) (*service.TimeReply, error) {
	now := time.Now().Unix()
	log.Printf("Received TimeRequest at %d", now)
	return &service.TimeReply{Timestamp: now}, nil
}

func main() {
	log.Printf("Starting server")

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	service.RegisterServiceServer(grpcServer, &serviceServer{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to start grpc: %v", err)
	}

}
