package note_v1

import (
	"context"

	"github.com/MaksMalf/testGrpc/internal/app/api/converter"
	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) GetListNote(ctx context.Context, in *emptypb.Empty) (*pb.GetListNoteResponce, error) {
	res, err := i.noteService.GetListNote(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetListNoteResponce{
		Notes: converter.ToPbNotes(res),
	}, nil
}
