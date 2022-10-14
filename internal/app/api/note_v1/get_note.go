package note_v1

import (
	"context"
	"fmt"

	pb "github.com/MaksMalf/test_gRPC/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, req *pb.GetNoteRequest) (*pb.GetNoteResponce, error) {
	fmt.Println("GetNote")
	fmt.Println("title:", req.GetId())

	return &pb.GetNoteResponce{
		Title:  "Hello!",
		Text:   "I'm create new request",
		Author: "Max",
	}, nil
}
