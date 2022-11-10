package storage

//go:generate mockgen --build_flags=--mod=mod -destination=mocks/mocks_note_storage.go -package=mocks . NoteStorage

import (
	"context"
	"time"

	"github.com/MaksMalf/testGrpc/internal/app/api/model"
	"github.com/MaksMalf/testGrpc/internal/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

const Note = "note"

type NoteStorage interface {
	CreateNote(ctx context.Context, note *model.NoteInfo) (int64, error)
	DeleteNote(ctx context.Context, noteID int64) error
	GetNote(ctx context.Context, noteID int64) (*model.Note, error)
	GetListNote(ctx context.Context) ([]*model.Note, error)
	UpdateNote(ctx context.Context, id int64, note *model.UpdateNoteInfo) error
}

type storage struct {
	client db.Client
}

func NewNoteStorage(client db.Client) NoteStorage {
	return &storage{client: client}
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

	q := db.Query{
		Name:     "CreateNote",
		QueryRaw: query,
	}

	row, err := s.client.DB().QueryContext(ctx, q, args...)
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

	q := db.Query{
		Name:     "DeleteNote",
		QueryRaw: query,
	}

	_, err = s.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) GetNote(ctx context.Context, noteID int64) (*model.Note, error) {
	builder := sq.Select("id, title, text, author, created_at, update_at").
		PlaceholderFormat(sq.Dollar).
		From(Note).
		Where(sq.Eq{"id": noteID}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetNote",
		QueryRaw: query,
	}

	var res model.Note
	err = s.client.DB().GetContext(ctx, &res, q, args...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *storage) GetListNote(ctx context.Context) ([]*model.Note, error) {
	builder := sq.Select("id, title, text, author, created_at, update_at").
		PlaceholderFormat(sq.Dollar).
		From(Note)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetListNote",
		QueryRaw: query,
	}

	var notes []*model.Note
	err = s.client.DB().SelectContext(ctx, &notes, q, args...)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (s *storage) UpdateNote(ctx context.Context, id int64, note *model.UpdateNoteInfo) error {
	builder := sq.Update(Note).
		PlaceholderFormat(sq.Dollar).
		Set("update_at", time.Now()).
		Set("title", note.Title.String).
		Set("text", note.Text.String).
		Set("author", note.Author.String).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "UpdateNote",
		QueryRaw: query,
	}

	_, err = s.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
