package main

import (
	"context"
	"net"
	"testing"

	pb "github.com/jeremyKisner/my-grcp-service/rollerService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestRoll(t *testing.T) {
	// Set up a gRPC server
	serverAddr := "localhost:50051" // You can use any available port
	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		t.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRollerServer(s, &server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			t.Fatalf("Server exited with error: %v", err)
		}
	}()
	defer s.Stop()

	// Set up a gRPC client to communicate with the server
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewRollerClient(conn)

	// Create a test request
	req := &pb.RollerRequest{Name: "TestRoll"}

	// Call the Roll function on the server
	resp, err := client.Roll(context.Background(), req)
	if err != nil {
		t.Fatalf("Roll failed: %v", err)
	}

	// Verify the response
	expectedResponse := &pb.RollerReply{Message: "Roller TestRoll"}
	if resp.Message != expectedResponse.Message {
		t.Errorf("Response was incorrect, got: %s, want: %s", resp.Message, expectedResponse.Message)
	}
}
