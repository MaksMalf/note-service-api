package note

import (
	"context"

	"github.com/MaksMalf/testGrpc/internal/app/api/model"
)

func (s *Service) UpdateNote(ctx context.Context, note *model.NoteInfo) error {
	return s.noteStorage.UpdateNote(ctx, note)
}
