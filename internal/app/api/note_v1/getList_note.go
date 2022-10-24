package note_v1

import (
	"context"
	"fmt"

	pb "github.com/MaksMalf/test_gRPC/pkg/note_v1"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

func (n *Note) GetListNote(ctx context.Context, req *pb.Empty) (*pb.GetListNoteResponce, error) {
	dbDsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbName, dbUser, dbPassword, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Select("id, title, text, author").
		PlaceholderFormat(sq.Dollar).
		From(noteTable)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var res []Note
	err = db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		return nil, err
	}

	notes := make([]*pb.GetListNoteResponce_Result_Note, 0, len(res))
	for _, n := range res {
		notes = append(notes, &pb.GetListNoteResponce_Result_Note{
			Id:     n.Id,
			Title:  n.Title,
			Text:   n.Text,
			Author: n.Author,
		})
	}

	return &pb.GetListNoteResponce{
		Result: &pb.GetListNoteResponce_Result{Notes: notes},
	}, nil
}
