package main

import (
	"context"
	"log"
	"time"

	pb "github.com/IbadT/go-lang-microservices.git/auth/pkg/auth"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
	userId  = 1
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect: ", err)
	}
	defer conn.Close()

	c := pb.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &pb.GetRequest{Id: userId})
	if err != nil {
		log.Fatal("failed to get user by id: ", err)
	}
	log.Printf("Note info:\n", color.GreenString("%+v", r.GetInfo()))
}
