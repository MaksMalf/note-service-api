package note_v1

import (
	"context"

	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
)

func (i *Implementation) GetNote(ctx context.Context, req *pb.GetNoteRequest) (*pb.GetNoteResponce, error) {
	res, err := i.noteService.GetNote(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetNoteResponce{
		Result: &pb.GetNoteResponce_Note{
			Id:     res.Id,
			Title:  res.Title,
			Text:   res.Text,
			Author: res.Author,
		},
	}, nil
}
