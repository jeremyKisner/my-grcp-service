package main

import (
	"context"
	"net"
	"testing"

	pb "github.com/jeremyKisner/my-grcp-service/rollerService"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Test_server_Roll(t *testing.T) {
	tests := []struct {
		name     string
		in       *pb.RollerRequest
		expected *pb.RollerReply
	}{
		{
			name:     "test 1",
			in:       &pb.RollerRequest{Name: "Roller"},
			expected: &pb.RollerReply{Message: "Roller"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up a gRPC server
			serverAddr := "localhost:50051"
			lis, err := net.Listen("tcp", serverAddr)
			assert.NoError(t, err)
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
			assert.NoError(t, err)
			defer conn.Close()

			client := pb.NewRollerClient(conn)
			ctx := context.Background()
			got, err := client.Roll(ctx, tt.in)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected.Message, got.Message)
			assert.NotNil(t, got.Total)
		})
	}
}
