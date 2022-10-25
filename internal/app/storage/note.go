package storage

import (
	"context"
	"errors"

	"github.com/MaksMalf/testGrpc/internal/storage/table"
	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type NoteStorage interface {
	CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (int64, error)
	DeleteNote(ctx context.Context, noteID int64) error
}

type storage struct {
	db *sqlx.DB
}

func NewNoteStorage(db *sqlx.DB) NoteStorage {
	return &storage{db: db}
}

func (s *storage) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (int64, error) {
	builder := sq.Insert(table.Note).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author").
		Values(req.GetTitle(), req.GetText(), req.GetAuthor()).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *storage) DeleteNote(ctx context.Context, noteID int64) error {
	builder := sq.Delete(table.Note).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": noteID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	row, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer row.Close()

	return nil
}

func (s *storage) GetNote(ctx context.Context, noteID int64) (*pb.GetNoteResponce, error) {
	builder := sq.Select("title, text, author").
		PlaceholderFormat(sq.Dollar).
		From(table.Note).
		Where(sq.Eq{"id": noteID}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var res []Implementation
	err = s.db.SelectContext(ctx, &res, query, args...)
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
