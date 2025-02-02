package main

import (
	"encoding/json"
	"fmt"

	"github.com/brianvoe/gofakeit"
	"github.com/brianvoe/gofakeit/v7"
	"google.golang.org/protobuf/proto"
)

func main() {
	session := &desc.NoteInfo{
		Title:    gofakeit.BeerName(),
		Content:  gofakeit.IPv4Address(),
		Author:   gofakeit.Name(),
		IsPublic: gofakeit.Bool(),
	}
	dataJson, _ := json.Marshal(session)
	// json занимает 91 byte
	fmt.Printf("\n\ndataJson len %d byte \n%v\n", len(dataJson), dataJson)

	dataPb, _ := proto.Marshal(session)
	// profobuf занимает 60 byte
	fmt.Printf("\n\ndataPb len %d byte \n%v\n", len(dataPb), dataPb)
}
