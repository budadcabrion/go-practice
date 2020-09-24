package main

import (
	"context"
	"log"
	"time"
	"os"

	"github.com/budadcabrion/go-practice/service"

	"google.golang.org/grpc"
)

const (
	address = "localhost:12345"
	defaultCmd = "time"
)

func main() {
	// connect to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	client := service.NewServiceClient(conn)

	// minimal arg parsing
	cmd := defaultCmd
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	// send a message to the server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch cmd {
	case "time":
		r, err := client.Time(ctx, &service.TimeRequest{})
		if err != nil {
			log.Fatalf("could not get time: %v", err)
		}
		t := time.Unix(r.Timestamp, 0)
		log.Printf("Time: %d, %s", r.Timestamp, t.String())

	// default:
	// 	log.Fatalf("invalid command: %v", cmd)
	}
}