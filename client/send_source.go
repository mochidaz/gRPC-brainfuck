package main

import (
	"log"
	pb "mochidaz/simple-service/proto"

	"context"
)

func sendSource(c pb.BrainfuckServiceClient) {
	log.Println("Invoked")
	res, err := c.Interpret(context.Background(), &pb.BrainfuckSourceRequest{Source: "+++++++++[>++++++++>+++++++++++>++++>+++++++++<<<<-]>.>++.+++++++..+++.>----.<<-.>>>+.--.<<<----.>>+."})

	if err != nil {
		log.Fatalf("could not interpret: %v", err)
	}

	log.Printf("Response from server: %s", res.Output)
}
