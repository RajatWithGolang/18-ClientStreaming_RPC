package main

import (
	"io"
	"log"
	"net"

	greetpb "github.com/Rajat2019/GRPC_IN_ACTION/03-ClientStreaming/proto"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(stream greetpb.GreetLongService_GreetServer) error {
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.GreetLongResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}

}
func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("this is an error %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetLongServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
