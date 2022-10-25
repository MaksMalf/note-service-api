package note

import (
	"context"

	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
)

func (s *Service) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (int64, error) {
	return s.noteStorage.CreateNote(ctx, req)
}
