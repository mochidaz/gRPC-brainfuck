package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "mochidaz/simple-service/proto"
)

type Server struct {
	pb.UnimplementedBrainfuckServiceServer
}

func main() {
	conn, err := net.Listen("tcp", "0.0.0.0:50051")

	log.Printf("Listening on %v\n", conn.Addr())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	log.Printf("Starting server")

	s := grpc.NewServer()

	pb.RegisterBrainfuckServiceServer(s, &Server{})

	log.Printf("Server started")
	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
