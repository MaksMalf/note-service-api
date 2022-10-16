package note_v1

import pb "github.com/MaksMalf/test_gRPC/pkg/note_v1"

type Note struct {
	Id     int64  `db:"id"`
	Title  string `db:"title"`
	Text   string `db:"text"`
	Author string `db:"author"`
	pb.UnimplementedNoteV1Server
}

func NewNote() *Note {
	return &Note{}
}
