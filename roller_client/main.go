package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eiannone/keyboard"
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
	Start()
}

func Start() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRollerClient(conn)

	err = keyboard.Open()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer keyboard.Close()

	fmt.Println("Press any key (Press 'q' or 'esc' to quit)")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println(err)
			break
		}
		if key == keyboard.KeyEsc || char == 'q' {
			log.Printf("terminating program")
			break
		} else if char == 'r' {
			// Contact the server and print out its response.
			ClientRoll(c)
		} else {
			fmt.Printf("unsupported key bind pressed: %s\n", string(char))
		}
	}
}

func ClientRoll(c pb.RollerClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Roll(ctx, &pb.RollerRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not roll: %v", err)
	}
	log.Printf("%s rolled a %s", r.GetMessage(), r.GetTotal())
}
