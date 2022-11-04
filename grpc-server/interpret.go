package grpc_server

import (
	"context"
	"log"
	pb "mochidaz/simple-service/proto"
	"mochidaz/simple-service/utils"
)

func (s *Server) Interpret(ctx context.Context, in *pb.BrainfuckSourceRequest) (*pb.OutputResponse, error) {
	log.Printf("Invoked %v", in)
	res, err := utils.Execute(in.Source)
	if err != nil {
		log.Fatalf("could not interpret: %v", err)
		return nil, err
	}
	return &pb.OutputResponse{Output: res}, nil
}
