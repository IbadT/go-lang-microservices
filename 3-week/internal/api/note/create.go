package note

import (
	"context"
	"log"

	"github.com/IbadT/go-lang-microservices.git/internal/converter"
	desc "github.com/IbadT/go-lang-microservices.git/pkg/note_v1"
	// "github.com/olezhek28/microservices_course/internal/converter"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.noteService.Create(ctx, converter.ToNoteInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted note with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
