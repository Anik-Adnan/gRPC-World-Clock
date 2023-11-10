package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Anik-Adnan/WorldClock/protoc"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTimeServiceClient(conn)

	req := &pb.TimeZoneRequest{
		Name: "America/New_York",
	}
	resp, err := client.GetTime(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to get time: %v", err)
	}

	fmt.Println("Current time in America/New_York:", resp.Time)
}
