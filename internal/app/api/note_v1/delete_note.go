package note_v1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
)

func (i *Implementation) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*emptypb.Empty, error) {
	res, err := i.noteService.DeleteNote(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return res, nil

}
