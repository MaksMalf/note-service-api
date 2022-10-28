package note

import (
	"context"

	"github.com/MaksMalf/testGrpc/internal/app/api/model"
)

func (s *Service) UpdateNote(ctx context.Context, id int64, note *model.UpdateNoteInfo) error {
	return s.noteStorage.UpdateNote(ctx, id, note)
}
