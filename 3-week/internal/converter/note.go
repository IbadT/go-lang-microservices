package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/IbadT/go-lang-microservices.git/internal/model"
	desc "github.com/IbadT/go-lang-microservices.git/pkg/note_v1"
)

func ToNoteFromService(note *model.Note) *desc.Note {
	var updatedAt *timestamppb.Timestamp
	if note.UpdatedAt.Valid {
		updatedAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return &desc.Note{
		Id:        note.ID,
		Info:      ToNoteInfoFromService(note.Info),
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToNoteInfoFromService(info model.NoteInfo) *desc.NoteInfo {
	return &desc.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
		// author
		// is_public
	}
}

func ToNoteInfoFromDesc(info *desc.NoteInfo) *model.NoteInfo {
	return &model.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}

//

func ToNoteListFromService(notes []model.Note) []*desc.Note {
	// Создаем срез для хранения преобразованных заметок
	var descNotes []*desc.Note

	// Проходим по каждой заметке в списке
	for _, note := range notes {
		// Преобразуем каждую заметку с помощью ToNoteFromService
		descNote := ToNoteFromService(&note)
		// Добавляем преобразованную заметку в срез
		descNotes = append(descNotes, descNote)
	}

	// Возвращаем срез преобразованных заметок
	return descNotes
}
