package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "mochidaz/simple-service/proto"
)

func sendSource(c pb.BrainfuckServiceClient) {
	log.Println("Invoked")
	res, err := c.Interpret(context.Background(), &pb.BrainfuckSourceRequest{Source: "+++++++++[>++++++++>+++++++++++>++++>+++++++++<<<<-]>.>++.+++++++..+++.>----.<<-.>>>+.--.<<<----.>>+."})

	if err != nil {
		log.Fatalf("could not interpret: %v", err)
	}

	log.Printf("Response from server: %s", res.Output)
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewBrainfuckServiceClient(conn)

	sendSource(c)
}
