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

	notes := make([]*pb.Note, 0, len(res))
	for _, n := range res {
		notes = append(notes, converter.ToPbNote(n))
	}

	return &pb.GetListNoteResponce{
		Notes: notes,
	}, nil
}
