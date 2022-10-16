package note_v1

import (
	"context"
	"fmt"

	pb "github.com/MaksMalf/test_gRPC/pkg/note_v1"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	noteTable  = "note"
	host       = "localhost"
	port       = "54321"
	dbName     = "note-service"
	dbUser     = "note-service-user"
	dbPassword = "note-service-password"
	sslMode    = "disable"
)

func (n *Note) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponce, error) {
	dbDsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbName, dbUser, dbPassword, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Insert(noteTable).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author").
		Values(req.GetTitle(), req.GetText(), req.GetAuthor()).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &pb.CreateNoteResponce{
		Id: id,
	}, nil
}
