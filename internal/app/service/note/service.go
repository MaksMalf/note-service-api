package note

import "github.com/MaksMalf/testGrpc/internal/app/storage"

type Service struct {
	noteStorage storage.NoteStorage
}

func NewService(noteStorage storage.NoteStorage) *Service {
	return &Service{noteStorage: noteStorage}
}
