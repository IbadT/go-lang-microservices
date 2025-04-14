package note

import (
	"context"

	desc "github.com/IbadT/go-lang-microservices.git/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	emptyValue, err := i.noteService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return emptyValue, nil
}
