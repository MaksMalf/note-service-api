package note

import (
	"context"

	"github.com/MaksMalf/testGrpc/internal/app/api/model"
)

func (s *Service) GetListNote(ctx context.Context) ([]*model.Note, error) {
	return s.noteStorage.GetListNote(ctx)
}
