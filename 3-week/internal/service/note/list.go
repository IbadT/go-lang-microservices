package note

import (
	"context"

	"github.com/IbadT/go-lang-microservices.git/internal/model"
)

func (s *serv) List(ctx context.Context, info *model.ListRequest) (*[]model.Note, error) {
	notes, err := s.noteRepository.List(ctx, info)
	if err != nil {
		return nil, err
	}

	return notes, nil
}
