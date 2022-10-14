package note_v1

import (
	"context"
	"fmt"
	desc "github.com/MaksMalf/test_gRPC/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponce, error) {
	fmt.Println("GetNote")
	fmt.Println("title:", req.GetId())

	return &desc.GetNoteResponce{
		Title:  "Hello!",
		Text:   "I'm create new request",
		Author: "Max",
	}, nil
}
