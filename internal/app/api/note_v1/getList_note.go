package note_v1

import (
	"context"
	"fmt"
	desc "github.com/MaksMalf/test_gRPC/pkg/note_v1"
)

func (n *Note) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponce, error) {
	fmt.Println("GetListNote")
	fmt.Println("title:", req.GetGetAll())

	return &desc.GetListNoteResponce{
		AllList: "List all note",
	}, nil
}
