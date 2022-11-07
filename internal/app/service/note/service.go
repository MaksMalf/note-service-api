package note

import "github.com/MaksMalf/testGrpc/internal/app/storage"

type Service struct {
	noteStorage storage.NoteStorage
}

func NewService(noteStorage storage.NoteStorage) *Service {
	return &Service{noteStorage: noteStorage}
}

func NewMockNoteService(deps ...any) *Service {
	is := Service{}

	for _, v := range deps {
		switch s := v.(type) {
		case storage.NoteStorage:
			is.noteStorage = s
		}
	}

	return &is
}
