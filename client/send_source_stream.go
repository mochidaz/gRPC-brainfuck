package main

import (
	"context"
	"fmt"
	"io"
	"log"
	pb "mochidaz/simple-service/proto"
)

func sendSourceStream(c pb.BrainfuckServiceClient) {
	log.Println("Invoked")

	stream, err := c.InterpretStream(context.Background())

	waitc := make(chan struct{})

	if err != nil {
		log.Fatalf("error while reading client stream: %v", err)
	}

	go func() {
		for {
			var source string
			_, err := fmt.Scanf("%s", &source)

			if err == io.EOF {
				close(waitc)
				return
			}

			if err != nil {
				log.Fatalf("error while reading client stream: %v", err)
			}

			err = stream.Send(&pb.BrainfuckSourceRequest{Source: source})

			if err != nil {
				log.Fatalf("error while sending data to client: %v", err)
			}
		}
	}()
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("error while reading client stream: %v", err)
				break
			}

			log.Println("Response from server: ", res.Output)
		}
		close(waitc)
	}()
	<-waitc
	return
}
