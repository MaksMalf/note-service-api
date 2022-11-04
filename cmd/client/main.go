package main

import (
	"context"
	"database/sql"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/MaksMalf/testGrpc/internal/app/api/converter"
	"github.com/MaksMalf/testGrpc/internal/app/api/model"
	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
)

const adress = "localhost:50051"

func main() {
	con, err := grpc.Dial(adress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := pb.NewNoteV1Client(con)

	//Create
	res, err := client.CreateNote(context.Background(), &pb.CreateNoteRequest{
		Info: converter.ToPbNoteInfo(&model.NoteInfo{
			Title:  "Ready",
			Text:   "Gogogogo",
			Author: "Maximus",
		}),
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("CreateNote ID:", res.GetId())

	//Get
	res1, err := client.GetNote(context.Background(), &pb.GetNoteRequest{
		Id: 15,
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("GetNote:\n", res1.GetNote())

	//Get all
	res2, err := client.GetListNote(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("GetListNote:\n", res2.GetNotes())

	// Update
	if _, err = client.UpdateNote(context.Background(), &pb.UpdateNoteRequest{
		Id: 16,
		UpdateInfo: converter.TpPbUpdateNoteInfo(&model.UpdateNoteInfo{
			Title:  sql.NullString{String: "Hey"},
			Text:   sql.NullString{String: "Update"},
			Author: sql.NullString{String: "New Maximus"},
		}),
	}); err != nil {
		log.Println(err.Error())
	}
	log.Println("Update note")

	//Delete
	if _, err = client.DeleteNote(context.Background(), &pb.DeleteNoteRequest{
		Id: 6,
	}); err != nil {
		log.Println(err.Error())
	}
	log.Println("Table deleted")
}
