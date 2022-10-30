package note_v1

import (
	"context"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/MaksMalf/testGrpc/internal/app/api/converter"
	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
)

func (i *Implementation) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponce, error) {
	id, err := i.noteService.CreateNote(ctx, converter.ToNoteInfo(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &pb.CreateNoteResponce{
		Id: id,
	}, nil
}
