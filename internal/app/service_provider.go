package app

import (
	"context"
	"log"

	"github.com/MaksMalf/testGrpc/internal/app/service/note"
	"github.com/MaksMalf/testGrpc/internal/app/storage"
	"github.com/MaksMalf/testGrpc/internal/config"
	"github.com/MaksMalf/testGrpc/internal/pkg/db"
)

type serviceProvider struct {
	db         db.Client
	configPath string
	config     config.Config

	noteStorage storage.NoteStorage

	noteService *note.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		connString, err := s.GetConfig().GetDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		dbc, err := db.NewClient(ctx, connString)
		if err != nil {
			log.Fatalf("cant't connect to db err: %s", err.Error())
		}

		s.db = dbc
	}

	return s.db
}

func (s *serviceProvider) GetConfig() config.Config {
	if s.config == nil {
		cfg, err := config.NewConfig(s.configPath)
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) GetNoteStorage(ctx context.Context) storage.NoteStorage {
	if s.noteStorage == nil {
		s.noteStorage = storage.NewNoteStorage(s.GetDB(ctx))
	}

	return s.noteStorage
}

func (s *serviceProvider) GetNoteService(ctx context.Context) *note.Service {
	if s.noteService == nil {
		s.noteService = note.NewService(s.GetNoteStorage(ctx))
	}

	return s.noteService
}
