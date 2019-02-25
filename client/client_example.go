package main

import (
	"context"
	"log"

	"github.com/major28/monitoring/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

func main() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := monitoring.NewMonitoringClient(conn)

	ctx := metadata.NewOutgoingContext(
		context.Background(),
		metadata.Pairs("username", "test", "password", "test"),
	)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	_, err = c.GetServiceInfoByName(ctx, &monitoring.GetServiceInfoByNameRequest{Name: "facebook.com"})
	if err != nil {
		log.Println("something went wrong", err)
	}

	_, err = c.GetHistoryRequests(ctx, &monitoring.EmptyRequest{})
	if err != nil {
		log.Println("something went wrong", err)
	}

	_, err = c.GetServiceWithMaxResponseTime(ctx, &monitoring.EmptyRequest{})
	if err != nil {
		log.Println("something went wrong", err)
	}

	_, err = c.GetServiceWithMinResponseTime(ctx, &monitoring.EmptyRequest{})
	if err != nil {
		log.Println("something went wrong", err)
	}
}
