package note

import (
	"context"

	"github.com/IbadT/go-lang-microservices.git/internal/converter"
	desc "github.com/IbadT/go-lang-microservices.git/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := i.noteService.Update(ctx, req.GetId(), converter.ToNoteInfoFromDesc(&desc.NoteInfo{
		Title:   req.Info.Title.GetValue(),
		Content: req.Info.Context.GetValue(),
	}))

	if err != nil {
		return &emptypb.Empty{}, nil
	}

	return &emptypb.Empty{}, nil
}
