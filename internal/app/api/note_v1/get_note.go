package note_v1

import (
	"context"

	"github.com/MaksMalf/testGrpc/internal/app/api/converter"
	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
)

func (i *Implementation) GetNote(ctx context.Context, req *pb.GetNoteRequest) (*pb.GetNoteResponce, error) {
	res, err := i.noteService.GetNote(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetNoteResponce{
		Note: converter.ToPbNote(res),
	}, nil
}
