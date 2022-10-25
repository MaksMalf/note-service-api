package main

import (
	"context"
	"log"

	pb "github.com/MaksMalf/testGrpc/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
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
		Title:  "Wow!",
		Text:   "I'm surprised",
		Author: "Max",
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("CreateNote ID:", res.GetResult().GetId())

	//Get
	res1, err := client.GetNote(context.Background(), &pb.GetNoteRequest{
		Id: 2,
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("GetNote:\n", res1.GetResult())

	//Get all
	res2, err := client.GetListNote(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("GetListNote:\n", res2.GetResult().GetNotes())

	// Update
	if _, err = client.UpdateNote(context.Background(), &pb.UpdateNoteRequest{
		Id:     6,
		Title:  "NewTitle",
		Text:   "NewText",
		Author: "NewMax",
	}); err != nil {
		log.Println(err.Error())
	}
	log.Println("Update note")

	// Delete
	if _, err = client.DeleteNote(context.Background(), &pb.DeleteNoteRequest{
		Id: 10,
	}); err != nil {
		log.Println(err.Error())
	}
	log.Println("Table deleted")
}
