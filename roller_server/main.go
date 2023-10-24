package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"

	pb "github.com/jeremyKisner/my-grcp-service/rollerService"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement Roller.RollerServer.
type server struct {
	pb.UnimplementedRollerServer
}

// Roll implements Roller.RollerServer
func (s *server) Roll(ctx context.Context, in *pb.RollerRequest) (*pb.RollerReply, error) {
	log.Printf("Received: %v", in.GetName())
	randomInt := rand.Intn(100) + 1
	log.Printf("Received: %v", randomInt)
	return &pb.RollerReply{
		Message: in.GetName(),
		Total:   fmt.Sprint(randomInt),
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRollerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
