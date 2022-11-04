package storage

import (
	"context"
	"errors"

	"github.com/MaksMalf/testGrpc/internal/app/api/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

const Note = "note"

type NoteStorage interface {
	CreateNote(ctx context.Context, note *model.NoteInfo) (int64, error)
	DeleteNote(ctx context.Context, noteID int64) error
	GetNote(ctx context.Context, noteID int64) (*model.NoteInfo, error)
	GetListNote(ctx context.Context) ([]*model.NoteInfo, error)
	UpdateNote(ctx context.Context, note *model.NoteInfo) error
}

type storage struct {
	db *sqlx.DB
}

func NewNoteStorage(db *sqlx.DB) NoteStorage {
	return &storage{db: db}
}

func (s *storage) CreateNote(ctx context.Context, note *model.NoteInfo) (int64, error) {
	builder := sq.Insert(Note).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author").
		Values(note.Title, note.Text, note.Author).
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
	builder := sq.Delete(Note).
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

func (s *storage) GetNote(ctx context.Context, noteID int64) (*model.NoteInfo, error) {
	builder := sq.Select("title, text, author").
		PlaceholderFormat(sq.Dollar).
		From(Note).
		Where(sq.Eq{"id": noteID}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var res []*model.NoteInfo
	err = s.db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		return nil, err
	}
	if len(res) <= 0 {
		return nil, errors.New("user not found")
	}

	return res[0], nil
}

func (s *storage) GetListNote(ctx context.Context) ([]*model.NoteInfo, error) {
	builder := sq.Select("id, title, text, author").
		PlaceholderFormat(sq.Dollar).
		From(Note)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var noteInfo []*model.NoteInfo
	err = s.db.SelectContext(ctx, &noteInfo, query, args...)
	if err != nil {
		return nil, err
	}

	return noteInfo, nil
}

func (s *storage) UpdateNote(ctx context.Context, note *model.NoteInfo) error {
	builder := sq.Update(Note).
		PlaceholderFormat(sq.Dollar).
		Set("title", note.Title).
		Set("text", note.Text).
		Set("author", note.Author).
		Where(sq.Eq{"id": note.Id})

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
