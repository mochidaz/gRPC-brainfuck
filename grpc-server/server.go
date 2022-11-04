package grpc_server

import pb "mochidaz/simple-service/proto"

type Server struct {
	pb.UnimplementedBrainfuckServiceServer
}
