package note_v1

import (
	"context"
	"fmt"

	pb "github.com/MaksMalf/test_gRPC/pkg/note_v1"
)

func (n *Note) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponce, error) {
	fmt.Println("CreateNote")
	fmt.Println("title:", req.GetTitle())
	fmt.Println("text:", req.GetText())
	fmt.Println("author:", req.GetAuthor())

	return &pb.CreateNoteResponce{
		Id: 1,
	}, nil
}
