package note_v1

import (
	"context"
	"fmt"
	desc "github.com/MaksMalf/test_gRPC/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponce, error) {
	fmt.Println("CreateNote")
	fmt.Println("title:", req.GetTitle())
	fmt.Println("text:", req.GetText())
	fmt.Println("author:", req.GetAuthor())

	return &desc.CreateNoteResponce{
		Id: 1,
	}, nil
}
