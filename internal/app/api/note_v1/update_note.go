package note_v1

import (
	"context"
	"fmt"
	desc "github.com/MaksMalf/test_gRPC/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponce, error) {
	fmt.Println("UpdateNote")
	fmt.Println("title:", req.GetTitle())

	return &desc.UpdateNoteResponce{
		NewTitle: "Title is update",
	}, nil
}
