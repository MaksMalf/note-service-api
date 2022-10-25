package note

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) DeleteNote(ctx context.Context, noteID int64) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.noteStorage.DeleteNote(ctx, noteID)

}
