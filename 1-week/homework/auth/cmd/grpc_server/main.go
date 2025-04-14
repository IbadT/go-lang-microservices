package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/IbadT/go-lang-microservices.git/auth/pkg/auth"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50051

type server struct {
	pb.UnimplementedAuthServer
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterAuthServer(s, &server{})

	log.Printf("server listening at %d", listen.Addr())
	if err = s.Serve(listen); err != nil {
		log.Fatalf("Failed to server: %+v", err)
	}
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("Note id: %d", req.Id)
	return &pb.GetResponse{
		Id: req.GetId(),
		Info: &pb.UserInfo{
			Email: gofakeit.Email(),
			Name:  gofakeit.Name(),
			Role:  pb.Role_ADMIN,
		},
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

// func (s *server) Register(ctx context.Context, req *pb.RegisterRequest) {

// }

// func (s *server) Update(ctx context.Context, req *pb.UpdateRequest) {

// }

// func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) {

// }
