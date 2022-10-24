package main

import (
	"fmt"
	"io"
	"log"
	pb "mochidaz/simple-service/proto"
	"mochidaz/simple-service/utils"
)

func (s *Server) InterpretStream(stream pb.BrainfuckService_InterpretStreamServer) error {
	log.Printf("Invoked %v", stream)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while reading client stream: %v", err)
			continue
		}

		bf, err := utils.Execute(req.Source)

		if err != nil {
			log.Fatalf("could not interpret: %v", err)
		}

		res := fmt.Sprintf("Resp: %s", bf)

		err = stream.Send(&pb.OutputResponse{Output: res})

		if err != nil {
			log.Fatalf("error while sending data to client: %v", err)
		}
	}
}
