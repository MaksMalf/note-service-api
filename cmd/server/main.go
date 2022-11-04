package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/MaksMalf/testGrpc/internal/app/api/note_v1"
	"github.com/MaksMalf/testGrpc/internal/app/service/note"
	"github.com/MaksMalf/testGrpc/internal/app/storage"
	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"

	"google.golang.org/grpc"
)

const (
	host       = "localhost"
	dbPort     = "54321"
	dbUser     = "note-service-user"
	dbPassword = "note-service-password"
	dbName     = "note-service"
	sslMode    = "disable"
	grpcPort   = ":50051"
	httpPort   = ":8000"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(startGRPC())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(startHTTP())
	}()

	wg.Wait()
}

func startGRPC() error {
	list, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("failed to mapping port: %s", err.Error())
	}
	defer list.Close()

	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, dbPort, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return fmt.Errorf("failed to open connection with db")
	}
	defer db.Close()

	noteStorage := storage.NewNoteStorage(db)
	noteService := note.NewService(noteStorage)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)
	pb.RegisterNoteV1Server(s, note_v1.NewNote(noteService))

	if err = s.Serve(list); err != nil {
		return fmt.Errorf("failed to server: %s", err.Error())
	}

	return nil
}

func startHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterNoteV1HandlerFromEndpoint(ctx, mux, grpcPort, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(httpPort, mux)
}
