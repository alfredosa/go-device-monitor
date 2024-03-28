package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/alfredosa/go-sensors-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedMonitorServiceServer
}

func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	fmt.Printf(values[0])
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	expectedToken := "Bearer " + os.Getenv("BEARER_TOKEN")
	if values[0] != expectedToken {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	// Proceed with the handler if the token is valid
	return handler(ctx, req)
}

func (s *server) SendUsage(ctx context.Context, in *pb.UsageData) (*pb.UsageResponse, error) {
	log.Printf("Received: CPU Usage: %v, RAM Usage: %v%% \n", in.CpuUsage, in.RamUsage.UsedPercent)

	log.Printf("Host %s, uptime %v", in.HostDetails.Hostname, in.HostDetails.Uptime)
	return &pb.UsageResponse{Message: "Usage Received"}, nil
}

func main() {

	port := os.Getenv("SERVER_PORT")

	if port == "" {
		log.Fatalf("SERVER_PORT must be set")
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor),
	)

	pb.RegisterMonitorServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
