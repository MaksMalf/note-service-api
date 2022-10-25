package note_v1

import (
	"github.com/MaksMalf/testGrpc/internal/app/service/note"
	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
)

type Implementation struct {
	pb.UnimplementedNoteV1Server
	noteService *note.Service
}

func NewNote(noteService *note.Service) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
