package note

import (
	"github.com/IbadT/go-lang-microservices.git/internal/service"
	desc "github.com/IbadT/go-lang-microservices.git/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedNoteV1Server
	noteService service.NoteService
}

func NewImplementation(noteService service.NoteService) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
