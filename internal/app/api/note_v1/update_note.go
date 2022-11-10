package note_v1

import (
	"context"

	"github.com/MaksMalf/testGrpc/internal/app/api/converter"
	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) UpdateNote(ctx context.Context, req *pb.UpdateNoteRequest) (*emptypb.Empty, error) {
	err := i.noteService.UpdateNote(ctx, req.GetId(), converter.ToUpdateNoteInfo(req.UpdateInfo))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
