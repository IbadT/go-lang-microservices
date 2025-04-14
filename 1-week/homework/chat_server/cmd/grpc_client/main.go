package main

import (
	"log"

	// pb "github.com/IbadT/go-lang-microservices.git/chat_server/pkg/chat_server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address  = "localhost:50051"
	deleteId = 1
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}
	defer conn.Close()

	// c := pb.NewChatServerClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	// r, err := c.SendMessage(ctx, &pb.DeleteRequest{Id: 12})
	// if err != nil {
	// 	log.Fatal("Failed to get User By Id: ", err)
	// }
	// log.Printf("Note info:\n", fmt.Sprintf(color.Green("%+v")))
}
