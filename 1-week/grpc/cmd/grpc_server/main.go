package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/IbadT/go-lang-microservices.git/pkg/note_v1"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50051

type server struct {
	// после этого страаивания, структура будет соответствовать всем методам
	desc.UnimplementedNoteV1Server
	// note_v1.UnimplementedNoteV1Server
}

// Get ...
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	// func (s *server) Get(ctx context.Context, req *note_v1.GetRequest) (*note_v1.GetResponse, error) {
	log.Printf("Note id: %d", req.Id)

	return &desc.GetResponse{
		Note: &desc.Note{
			Id: req.GetId(),
			Info: &desc.NoteInfo{
				Title:    gofakeit.BeerName(),
				Content:  gofakeit.IPv4Address(),
				Author:   gofakeit.Name(),
				IsPublic: gofakeit.Bool(),
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	s := grpc.NewServer()
	// для того, чтобы получать информацию из сервера
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, &server{})
	// note_v1.RegisterNoteV1Server(s, &server{})

	log.Printf("server listening at %d", listen.Addr())

	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to server: #{err}")
	}
}
