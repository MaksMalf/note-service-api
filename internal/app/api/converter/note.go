package converter

import (
	"database/sql"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/MaksMalf/testGrpc/internal/app/api/model"
	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
)

func ToNoteInfo(noteInfo *pb.NoteInfo) *model.NoteInfo {
	return &model.NoteInfo{
		Title:  noteInfo.GetTitle(),
		Text:   noteInfo.GetText(),
		Author: noteInfo.GetAuthor(),
	}
}

func ToPbNoteInfo(noteInfo *model.NoteInfo) *pb.NoteInfo {
	return &pb.NoteInfo{
		Title:  noteInfo.Title,
		Text:   noteInfo.Text,
		Author: noteInfo.Author,
	}
}

func ToPbNote(note *model.Note) *pb.Note {
	var updateAt *timestamppb.Timestamp
	if note.UpdateAt.Valid {
		updateAt = timestamppb.New(note.UpdateAt.Time)
	}

	return &pb.Note{
		Id:        note.ID,
		Info:      ToPbNoteInfo(note.Info),
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdateAt:  updateAt,
	}
}

func ToPbNotes(res []*model.Note) []*pb.Note {
	notes := make([]*pb.Note, 0, len(res))
	for _, n := range res {
		notes = append(notes, ToPbNote(n))
	}

	return notes
}

func ToUpdateNoteInfo(updateInfo *pb.UpdateNoteInfo) *model.UpdateNoteInfo {
	return &model.UpdateNoteInfo{
		Title: sql.NullString{
			String: updateInfo.GetTitle().GetValue(),
			Valid:  updateInfo.GetTitle() != nil,
		},
		Text: sql.NullString{
			String: updateInfo.GetText().GetValue(),
			Valid:  updateInfo.GetText() != nil,
		},
		Author: sql.NullString{
			String: updateInfo.GetAuthor().GetValue(),
			Valid:  updateInfo.GetAuthor() != nil,
		},
	}
}

func TpPbUpdateNoteInfo(updateInfo *model.UpdateNoteInfo) *pb.UpdateNoteInfo {
	return &pb.UpdateNoteInfo{
		Title:  wrapperspb.String(updateInfo.Title.String),
		Text:   wrapperspb.String(updateInfo.Text.String),
		Author: wrapperspb.String(updateInfo.Author.String),
	}
}
