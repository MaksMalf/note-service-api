package note_v1

import (
	"context"
	"fmt"

	pb "github.com/MaksMalf/test_gRPC/pkg/note_v1"
	//sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

func (n *Note) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*pb.Empty, error) {
	dbDsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbName, dbUser, dbPassword, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	//builder := sq.Delete(noteTable).
	//	PlaceholderFormat(sq.Dollar).
	//	Where(sq.Eq{"id": req.GetId()})
	//
	//query, args, err := builder.ToSql()
	//if err != nil {
	//	return nil, err
	//}

	row, err := db.QueryContext(ctx, "DELETE FROM note WHERE id = $1", req.GetId())
	if err != nil {
		return nil, err
	}
	defer row.Close()

	return &pb.Empty{}, nil
}
