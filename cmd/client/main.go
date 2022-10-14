package main

import (
	"context"
	desc "github.com/MaksMalf/test_gRPC/pkg/note_v1"
	"google.golang.org/grpc"
	"log"
)

const adress = "localhost:50051"

func main() {
	con, err := grpc.Dial(adress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteV1Client(con)

	res, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "Wow!",
		Text:   "I'm surprised",
		Author: "Max",
	})

	res1, err := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Id: 1,
	})

	res2, err := client.GetListNote(context.Background(), &desc.GetListNoteRequest{
		GetAll: "Get list all",
	})

	res3, err := client.UpdateNote(context.Background(), &desc.UpdateNoteRequest{
		Id:    1,
		Title: "New Title",
	})

	res4, err := client.DeleteNote(context.Background(), &desc.DeleteNoteRequest{
		Id: 1,
	})

	if err != nil {
		log.Println(err.Error())
	}

	log.Println("CreateNote ID:", res.GetId())
	log.Println("GetNote Author:", res1.GetAuthor())
	log.Println("GetListNote Title:", res2.GetAllList())
	log.Println("UpdateNote Title:", res3.GetNewTitle())
	log.Println("DeleteNote Message:", res4.GetDelNote())
}
