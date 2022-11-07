package note

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/MaksMalf/testGrpc/internal/app/api/model"
)

func (s *Service) GetNote(ctx context.Context, noteID int64) (*model.Note, error) {
	note, err := s.noteStorage.GetNote(ctx, noteID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "note with this id is missing")
		}
		return nil, err
	}
	
	return note, nil
}
