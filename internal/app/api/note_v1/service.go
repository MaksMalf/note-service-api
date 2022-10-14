package note_v1

import pb "github.com/MaksMalf/test_gRPC/pkg/note_v1"

type Note struct {
	pb.UnimplementedNoteV1Server
}

func NewNote() *Note {
	return &Note{}
}
