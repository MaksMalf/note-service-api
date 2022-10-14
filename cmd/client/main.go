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

	res, err := client.CreateNote(context.Background(), &pb.CreateNoteRequest{
		Title:  "Wow!",
		Text:   "I'm surprised",
		Author: "Max",
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("CreateNote ID:", res.GetId())

	res1, err := client.GetNote(context.Background(), &pb.GetNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("GetNote Author:", res1.GetAuthor())

	res2, err := client.GetListNote(context.Background(), &pb.GetListNoteRequest{
		GetAll: "Get list all",
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("GetListNote Title:", res2.GetAllList())

	res3, err := client.UpdateNote(context.Background(), &pb.UpdateNoteRequest{
		Id:    1,
		Title: "New Title",
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("UpdateNote Title:", res3.GetNewTitle())

	res4, err := client.DeleteNote(context.Background(), &pb.DeleteNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("DeleteNote Message:", res4.GetDelNote())
}
