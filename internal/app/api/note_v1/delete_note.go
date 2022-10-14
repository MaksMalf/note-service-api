package note_v1

import (
	"context"
	"fmt"
	desc "github.com/MaksMalf/test_gRPC/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.DeleteNoteResponce, error) {
	fmt.Println("DeleteNote")
	fmt.Println("Id:", req.GetId())

	return &desc.DeleteNoteResponce{
		DelNote: "Note is deleted",
	}, nil
}
