package note_v1

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/MaksMalf/test_gRPC/pkg/note_v1"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

func (n *Note) GetNote(ctx context.Context, req *pb.GetNoteRequest) (*pb.GetNoteResponce, error) {
	dbDsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbName, dbUser, dbPassword, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Select("title, text, author").
		PlaceholderFormat(sq.Dollar).
		From(noteTable).
		Where(sq.Eq{"id": req.GetId()}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var res []Note
	err = db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		return nil, err
	}
	if len(res) <= 0 {
		return nil, errors.New("user not found")
	}

	return &pb.GetNoteResponce{
		Title:  res[0].Title,
		Text:   res[0].Text,
		Author: res[0].Author,
	}, nil
}
