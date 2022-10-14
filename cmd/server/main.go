package main

import (
	"fmt"
	"log"
	"net"

	"github.com/MaksMalf/test_gRPC/internal/app/api/note_v1"
	pb "github.com/MaksMalf/test_gRPC/pkg/note_v1"

	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterNoteV1Server(s, note_v1.NewNote())

	fmt.Println("server run on port: %s", port)

	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}

}
