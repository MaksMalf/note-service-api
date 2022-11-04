package note

import (
	"context"

	"github.com/MaksMalf/testGrpc/internal/app/api/model"
)

func (s *Service) CreateNote(ctx context.Context, note *model.NoteInfo) (int64, error) {
	return s.noteStorage.CreateNote(ctx, note)
}
