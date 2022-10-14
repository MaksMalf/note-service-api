package note_v1

import (
	"context"
	"fmt"

	pb "github.com/MaksMalf/test_gRPC/pkg/note_v1"
)

func (n *Note) GetListNote(ctx context.Context, req *pb.GetListNoteRequest) (*pb.GetListNoteResponce, error) {
	fmt.Println("GetListNote")
	fmt.Println("title:", req.GetGetAll())

	return &pb.GetListNoteResponce{
		AllList: "List all note",
	}, nil
}
