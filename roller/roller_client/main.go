package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eiannone/keyboard"
	pb "github.com/jeremyKisner/my-grcp-service/roller/rollerService"
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
		log.Printf("error opening keyboard: %v:", err)
		os.Exit(1)
	}
	defer keyboard.Close()

	log.Printf("Press any key (Press 'q' or 'esc' to quit)")
	sessionWins := 0
	sessionLosses := 0
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Print(err)
			break
		}
		if key == keyboard.KeyEsc || char == 'q' {
			log.Printf("%s Wins - %s Losses", fmt.Sprint(sessionWins), fmt.Sprint(sessionLosses))
			log.Printf("game over")
			break
		} else if char == 'r' {
			// Contact the server and print out its response.
			ClientRoll(c, &sessionWins, &sessionLosses)
		} else {
			log.Printf("unsupported key bind pressed: %s\n", string(char))
		}
	}
}

func ClientRoll(c pb.RollerClient, wins *int, losses *int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Roll(ctx, &pb.RollerRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not roll: %v", err)
	}
	if r.GetPlayerWins() {
		*wins++
		log.Printf("You Win! %s rolled a %s to win over %s!", r.GetMessage(), fmt.Sprint(r.GetTotal()), fmt.Sprint(r.GetAiScore()))
	} else {
		*losses++
		log.Printf("You Lose! %s rolled a %s to lose against %s!", r.GetMessage(), fmt.Sprint(r.GetTotal()), fmt.Sprint(r.GetAiScore()))
	}
}
