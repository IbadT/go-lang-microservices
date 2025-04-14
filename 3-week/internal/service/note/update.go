package note

import (
	"context"

	"github.com/IbadT/go-lang-microservices.git/internal/model"
)

func (s *serv) Update(ctx context.Context, id int64, info *model.NoteInfo) error {
	err := s.noteRepository.Update(ctx, id, info)
	if err != nil {
		return err
	}

	return nil
}
