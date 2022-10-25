package main

import (
	"context"
	"log"

	pb "github.com/MaksMalf/test_gRPC/pkg/note_v1"

	"google.golang.org/grpc"
)

const adress = "localhost:50051"

func main() {
	con, err := grpc.Dial(adress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := pb.NewNoteV1Client(con)

	// Create
	res, err := client.CreateNote(context.Background(), &pb.CreateNoteRequest{
		Title:  "Wow!",
		Text:   "I'm surprised",
		Author: "Max",
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("CreateNote ID:", res.GetId())

	// Get
	res1, err := client.GetNote(context.Background(), &pb.GetNoteRequest{
		Id: 6,
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("GetNote Author:", res1.GetAuthor())

	//Get all
	res2, err := client.GetListNote(context.Background(), &pb.Empty{})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("GetListNote Title:", res2.GetResult())

	// Update
	res3, err := client.UpdateNote(context.Background(), &pb.UpdateNoteRequest{
		Id:     2,
		Title:  "NewTitle",
		Text:   "NewText",
		Author: "NewMax",
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(res3.GetNewTitle())

	// Delete
	if _, err := client.DeleteNote(context.Background(), &pb.DeleteNoteRequest{Id: 8}); err != nil {
		log.Println(err.Error())
	}
	log.Println("Table deleted")
}
