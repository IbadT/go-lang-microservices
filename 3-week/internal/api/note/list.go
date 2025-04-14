package note

import (
	"context"

	// "github.com/IbadT/go-lang-microservices.git/internal/model"

	"github.com/IbadT/go-lang-microservices.git/internal/converter"
	"github.com/IbadT/go-lang-microservices.git/internal/model"
	desc "github.com/IbadT/go-lang-microservices.git/pkg/note_v1"
)

func (i *Implementation) List(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
	notesObj, err := i.noteService.List(ctx, &model.ListRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, err
	}

	return &desc.ListResponse{
		Notes: converter.ToNoteListFromService(*notesObj),
	}, nil
}
