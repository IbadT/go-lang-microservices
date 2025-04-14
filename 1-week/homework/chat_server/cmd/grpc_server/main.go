package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/IbadT/go-lang-microservices.git/chat_server/pkg/chat_server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

type server struct {
	pb.UnimplementedChatServerServer
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterChatServerServer(s, server{})

	log.Printf("Server listening at %d", listen.Addr())
	if err = s.Serve(listen); err != nil {
		log.Fatalf("Failed to server: %+v", err)
	}
}
