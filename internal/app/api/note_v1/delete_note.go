package note_v1

import (
	"context"
	"fmt"

	pb "github.com/MaksMalf/test_gRPC/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*pb.DeleteNoteResponce, error) {
	fmt.Println("DeleteNote")
	fmt.Println("Id:", req.GetId())

	return &pb.DeleteNoteResponce{
		DelNote: "Note is deleted",
	}, nil
}
