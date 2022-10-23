package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"mochidaz/simple-service/utils"
	"net"

	pb "mochidaz/simple-service/proto"
)

type Server struct {
	pb.UnimplementedBrainfuckServiceServer
}

func (s *Server) Interpret(ctx context.Context, in *pb.BrainfuckSourceRequest) (*pb.OutputResponse, error) {
	log.Printf("Invoked %v", in)
	res, err := utils.Execute(in.Source)
	if err != nil {
		log.Fatalf("could not interpret: %v", err)
		return nil, err
	}
	return &pb.OutputResponse{Output: fmt.Sprintf("Resp: %s", res)}, nil
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
