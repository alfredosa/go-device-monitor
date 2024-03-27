package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/alfredosa/go-sensors-server/proto"
	"google.golang.org/grpc"
	// Update this import path
)

type server struct {
	pb.UnimplementedMonitorServiceServer
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

	s := grpc.NewServer()

	pb.RegisterMonitorServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
