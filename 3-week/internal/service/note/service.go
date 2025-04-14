package note

import (
	"github.com/IbadT/go-lang-microservices.git/internal/client/db"
	"github.com/IbadT/go-lang-microservices.git/internal/repository"
	"github.com/IbadT/go-lang-microservices.git/internal/service"
)

type serv struct {
	noteRepository repository.NoteRepository
	txManager      db.TxManager
}

func NewService(
	noteRepository repository.NoteRepository,
	txManager db.TxManager,
) service.NoteService {
	return &serv{
		noteRepository: noteRepository,
		txManager:      txManager,
	}
}
