package note_v1

import (
	"context"
	"fmt"

	pb "github.com/MaksMalf/test_gRPC/pkg/note_v1"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

func (n *Note) UpdateNote(ctx context.Context, req *pb.UpdateNoteRequest) (*pb.UpdateNoteResponce, error) {
	dbDsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbName, dbUser, dbPassword, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Update(noteTable).
		PlaceholderFormat(sq.Dollar).
		Set("title", req.GetTitle()).
		Set("text", req.GetText()).
		Set("author", req.GetAuthor()).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	return &pb.UpdateNoteResponce{
		NewTitle: "Note is update",
	}, nil
}
