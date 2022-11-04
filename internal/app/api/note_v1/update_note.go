package note_v1

import (
	"context"

	"github.com/MaksMalf/testGrpc/internal/app/api/model"
	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) UpdateNote(ctx context.Context, req *pb.UpdateNoteRequest) (*emptypb.Empty, error) {
	err := i.noteService.UpdateNote(ctx, &model.NoteInfo{
		Id:     req.GetId(),
		Title:  req.GetTitle(),
		Text:   req.GetText(),
		Author: req.GetAuthor(),
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
