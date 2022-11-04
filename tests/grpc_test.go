package tests

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	server "mochidaz/simple-service/grpc-server"
	pb "mochidaz/simple-service/proto"
	"net"
	"testing"
)

func initServer() {
	conn, err := net.Listen("tcp", "0.0.0.0:50051")

	log.Printf("Listening on %v\n", conn.Addr())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	log.Printf("Starting server")

	s := grpc.NewServer()

	pb.RegisterBrainfuckServiceServer(s, &server.Server{})

	log.Printf("Server started")
	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initClient() pb.BrainfuckServiceClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return pb.NewBrainfuckServiceClient(conn)
}

func sendSource(c pb.BrainfuckServiceClient, message string) string {
	log.Println("Invoked")
	res, err := c.Interpret(context.Background(), &pb.BrainfuckSourceRequest{Source: message})

	if err != nil {
		log.Fatalf("could not interpret: %v", err)
	}

	return res.Output
}

func TestUnary(t *testing.T) {

	go initServer()

	c := initClient()
	out_1 := sendSource(c, "+++++++++[>++++++++>+++++++++++>++++>++++++++++<<<<-]>.>++.+++++++..+++.>----.>---.<<.+++.------.--------.>+.")
	out_2 := sendSource(c, "+++++++++[>+++++++++>++++++++++++>++++>++++++++<<<<-]>+++.>----.+.++++++++++.>----.<----------.++++++++++.>.>-.<<<--.--.>>>----.<+.")
	out_3 := sendSource(c, "+++++++++[>++++++++++>+++++++++++>++++>+++++++++++++>++++++++<<<<<-]>---.>++.>----.<----.>>---.<<++++.>.>++.<<.>>-.+.<<++++.+++++.-------.>.>.<<+.---.>.>>-.<<<<-----.--.>>>>----.<<+.")

	if out_1 != "Hello World!" {
		panic("Test Failed!")
	}

	if out_2 != "This is GRPC!" {
		panic("Test Failed!")
	}

	if out_3 != "We are testing the GRPC!" {
		panic("Test Failed!")
	}

}
