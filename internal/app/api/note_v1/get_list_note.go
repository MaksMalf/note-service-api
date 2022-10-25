package note_v1

import (
	"context"

	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) GetListNote(ctx context.Context, in *emptypb.Empty) (*pb.GetListNoteResponce, error) {
	noteInfo, err := i.noteService.GetListNote(ctx)
	if err != nil {
		return nil, err
	}

	notes := make([]*pb.GetListNoteResponce_Result_Note, 0, len(noteInfo))
	for _, n := range noteInfo {
		notes = append(notes, &pb.GetListNoteResponce_Result_Note{
			Id:     n.Id,
			Title:  n.Title,
			Text:   n.Text,
			Author: n.Author,
		})
	}

	return &pb.GetListNoteResponce{
		Result: &pb.GetListNoteResponce_Result{
			Notes: notes,
		},
	}, nil
}
