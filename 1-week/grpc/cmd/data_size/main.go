package main

import (
	"encoding/json"
	"fmt"

	// desc "github.com/IbadT/go-lang-microservices/week_1/grpc/pkg/note_v1"

	desc "github.com/IbadT/go-lang-microservices.git/pkg/note_v1"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/protobuf/proto"
)

func main() {
	// session := &desc.NoteInfo{
	session := &desc.NoteInfo{
		Title:    gofakeit.BeerName(),
		Content:  gofakeit.IPv4Address(),
		Author:   gofakeit.Name(),
		IsPublic: gofakeit.Bool(),
	}

	dataJson, _ := json.Marshal(session)
	// json занимает 93 byte
	fmt.Printf("\n\ndataJson len %d byte \n%v\n", len(dataJson), dataJson)

	dataPb, _ := proto.Marshal(session)
	// profobuf занимает 62 byte
	fmt.Printf("\n\ndataPb len %d byte \n%v\n", len(dataPb), dataPb)
}
