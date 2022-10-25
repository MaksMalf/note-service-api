package note_v1

import (
	"context"

	"github.com/MaksMalf/testGrpc/internal/app/api/model"
	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
	_ "github.com/jackc/pgx/stdlib"
)

func (i *Implementation) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponce, error) {
	id, err := i.noteService.CreateNote(ctx, &model.NoteInfo{
		Title:  req.GetTitle(),
		Text:   req.GetText(),
		Author: req.GetAuthor(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateNoteResponce{
		Result: &pb.CreateNoteResponce_Result{
			Id: id,
		},
	}, nil
}
