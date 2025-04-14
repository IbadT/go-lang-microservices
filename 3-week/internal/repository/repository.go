package repository

import (
	"context"

	"github.com/IbadT/go-lang-microservices.git/internal/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NoteRepository interface {
	Create(ctx context.Context, info *model.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Note, error)
	List(ctx context.Context, info *model.ListRequest) (*[]model.Note, error)
	Update(ctx context.Context, id int64, info *model.NoteInfo) error
	Delete(ctx context.Context, id int64) (*emptypb.Empty, error)
}
