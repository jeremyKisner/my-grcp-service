package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/jeremyKisner/my-grcp-service/rollerService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "player_name"
)

var (
	addr = flag.String("addr", "localhost:50051", "address:port to connect to")
	name = flag.String("name", defaultName, "Name to roll for")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRollerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Roll(ctx, &pb.RollerRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not roll: %v", err)
	}
	log.Printf("%s rolled a %s", r.GetMessage(), r.GetTotal())
}
