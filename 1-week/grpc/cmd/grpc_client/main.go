package main

import (
	"context"
	"log"
	"time"

	desc "github.com/IbadT/go-lang-microservices.git/pkg/note_v1"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
	noteID  = 12
)

func main() {
	// указываем, что подключение не секьюрное - grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	// обращаемся к нашему сгенерированному пакету и создаем нового клиента, передовая туда connection
	c := desc.NewNoteV1Client(conn)

	// задаем timeout
	// если запрос будет больше 1 секунды, то мы вывалимся по timeout(если сервер будет не отвечать дольше 1 сек)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: noteID})
	if err != nil {
		log.Fatal("failed to get note by id: ", err)
	}

	log.Printf("Note info:\n", color.GreenString("%+v", r.GetNote()))
}
