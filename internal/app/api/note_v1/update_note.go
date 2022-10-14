package note_v1

import (
	"context"
	"fmt"

	pb "github.com/MaksMalf/test_gRPC/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *pb.UpdateNoteRequest) (*pb.UpdateNoteResponce, error) {
	fmt.Println("UpdateNote")
	fmt.Println("title:", req.GetTitle())

	return &pb.UpdateNoteResponce{
		NewTitle: "Title is update",
	}, nil
}
