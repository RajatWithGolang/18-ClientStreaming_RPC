package main

import (
	"context"
	"fmt"
	"log"
	"time"

	greetpb "github.com/Rajat2019/GRPC_IN_ACTION/03-ClientStreaming/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetLongServiceClient(conn)

	doClientStreaming(c)

}

func doClientStreaming(c greetpb.GreetLongServiceClient) {

	requests := []*greetpb.GreetLongRequest{
		&greetpb.GreetLongRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Stephane",
			},
		},
		&greetpb.GreetLongRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
			},
		},
		&greetpb.GreetLongRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Lucy",
			},
		},
		&greetpb.GreetLongRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mark",
			},
		},
		&greetpb.GreetLongRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Piper",
			},
		},
	}
	//fmt.Printf("Created Client %f", c)
	stream, err := c.Greet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling rpc: %v", err)
	}
	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)
}
